#include <WiFi.h>
#include <HTTPClient.h>
#include "ArduinoJson.h"
#include "SPIFFS.h"
#include <TimeLib.h>
#include <ArduinoJWT.h>
#include <sha256.h>

ArduinoJWT jwt = ArduinoJWT("bakuretsu");

#define CONNECT 0
#define REGISTRATION 1
#define QUESTION 2
#define SEND_DATA 3
#define RESET 4
#define SYNC 5
#define HISTORY 6

const char* ssid = "";
const char* password =  "";
boolean ssidFlag = true;
boolean passFlag = false;

boolean credFlag = false;
String credential = "";
boolean sensorFlag = false;
String sensorName = "";

static int state = 0;
static int lastState = 0;
String STATE = "";
String inputSerial;

int period = 10000;
unsigned long time_now = 0;

String id;
String regTime;

boolean regFlag = false;
boolean hisFlag = false;
int n = 0;

long  sensorVal;

tmElements_t tm;
int Year, Month, Day, Hour, Minute, Second ;

int count = 0;

void setup() {
  // put your setup code here, to run once:
  Serial.begin(115200);
  if (!SPIFFS.begin(true)) {
    Serial.println("An Error has occurred while mounting SPIFFS");
    return;
  }
}

void loop() {
  // put your main code here, to run repeatedly:
  while (Serial.available() && (state != SEND_DATA)) {
    inputSerial = Serial.readString(); // read the incoming data as string
    if (inputSerial == "RESET") {
      state = RESET;
    } else if (inputSerial == "SYNC") {
      state = SYNC;
    } else if (inputSerial == "HISTORY") {
      hisFlag = true;
      state = HISTORY;
    } else if (inputSerial == "DATA") {
      File fileToRead = SPIFFS.open("/test.txt");
      if (!fileToRead) {
        Serial.println("Failed to open file for reading");
        return;
      }
      Serial.println("Data:");
      while (fileToRead.available()) {
        Serial.write(fileToRead.read());
      }
      fileToRead.close();

    } else {
      if ((state == CONNECT) && (ssidFlag)) {
        ssid = inputSerial.c_str();
        ssidFlag = false;
        passFlag = true;
        sensorFlag = false;
      } else if ((state == CONNECT) && (passFlag)) {
        password = inputSerial.c_str();
        ssidFlag = false;
        passFlag = false;
        credFlag = true;
      } else if ((state == REGISTRATION) && (credFlag)) {
        credential = inputSerial;
        credFlag = false;
        sensorFlag = true;
      } else if ((state == REGISTRATION) && (sensorFlag)) {
        sensorName = inputSerial;
        sensorFlag = false;
      } else if ((state == HISTORY) && (hisFlag)) {
        hisFlag = false;
        n = inputSerial.toInt();
      } else {
        inputSerial = inputSerial + " not recognize";
      }
    }

    Serial.print("INPUT SERIAL : ");
    Serial.println(inputSerial);
  }
  lastState = state;
  fsm();
  if (lastState != state) {
    switch (state) {
      case CONNECT:
        STATE = "CONNECT";
        break;
      case REGISTRATION:
        STATE = "REGISTRATION";
        break;
      case QUESTION:
        STATE = "QUESTION";
        break;
      case SEND_DATA:
        STATE = "SEND_DATA";
        break;
      case RESET:
        STATE = "RESET";
        break;
      case SYNC:
        STATE = "SYNC";
        break;
      case HISTORY:
        STATE = "HISTORY";
        break;
    }

    Serial.print("STATE : ");
    Serial.print(STATE);
    Serial.println();
  }
  delay(1000);

  sensorVal = random(100);
}

void fsm() {
  switch (state) {
    case CONNECT:
      if (ssidFlag) {
        Serial.println("Please input ssid!!!");
      }
      if (passFlag) {
        Serial.println("Please input password!!!");
      }

      if (!ssidFlag && !passFlag) {

        WiFi.begin(ssid, password);

        while (WiFi.status() != WL_CONNECTED) { //Check for the connection
          delay(1000);
          Serial.println("Connecting to WiFi..");
        }

        state = REGISTRATION;
      }
      break;
    case REGISTRATION:
      if (credFlag) {
        Serial.println("Please input credential!!!");
      }
      if (sensorFlag) {
        Serial.println("Please input sensor name!!!");
      }

      if (!credFlag && !sensorFlag) {
        Serial.print("Your sensor name is ");
        Serial.println(sensorName);
        if ((WiFi.status() == WL_CONNECTED)) { //Check the current connection status

          HTTPClient http;

          http.begin("http://192.168.137.1:8080/registration"); //Specify the URL
          http.addHeader("Content-Type", "text/plain"); //Specify content-type header

          String message = "{\"credential_pass\":\"" + credential + "\"}";
          String encoded = jwt.encodeJWT(message);;
          Serial.println(encoded);
          int httpCode = http.POST(encoded); //Make the request

          if (httpCode > 0) { //Check for the returning code

            String payload = http.getString();

            String decoded = "";
            jwt.decodeJWT(payload, decoded);

            const size_t capacityReg = JSON_OBJECT_SIZE(2) + 40;
            DynamicJsonBuffer jsonBuffer(capacityReg);
            JsonObject& rootReg = jsonBuffer.parseObject(decoded);

            id = rootReg["id"].as<String>(); // "24_QCDL"
            regTime = rootReg["time"].as<String>(); // "21-04-2020 17:26:46"

            Serial.println("RESPONSE : ");
            //          Serial.println(httpCode);
            Serial.print("id : ");
            Serial.println(id);
            Serial.print("time : ");
            Serial.println(regTime);

            char dateTime[18] = {};
            regTime.toCharArray(dateTime, 18);
            createElements(dateTime);
            setTime(makeTime(tm));//set Ardino system clock to compiled time

            File fileToWrite = SPIFFS.open("/test.txt", FILE_WRITE);

            if (!fileToWrite) {
              Serial.println("There was an error opening the file for writing");
              return;
            }

            if (fileToWrite.println("id : " + id + ", " + "registration time : " + regTime + ", " + "sensor name : " + sensorName)) {
              Serial.println("File was written");
            } else {
              Serial.println("File write failed");
            }
            fileToWrite.close();
            state = QUESTION;
          } else {
            state = CONNECT;
            ssidFlag = true;
            Serial.println("Not Connected or credential false");
          }

          http.end(); //Free the resources
        }
      }
      break;
    case QUESTION:
      if ((WiFi.status() == WL_CONNECTED)) { //Check the current connection status

        HTTPClient http;

        http.begin("http://192.168.137.1:8080/question/" + id); //Specify the URL
        int httpCode = http.GET();                                        //Make the request

        if (httpCode > 0) { //Check for the returning code

          String payload = http.getString();

          String decoded = "";
          jwt.decodeJWT(payload, decoded);

          const size_t capacityQue = JSON_OBJECT_SIZE(1) + 20;
          DynamicJsonBuffer jsonBuffer(capacityQue);
          JsonObject& rootQue = jsonBuffer.parseObject(decoded);

          String stat = rootQue["status"].as<String>(); //

          Serial.println("RESPONSE : ");
          //          Serial.println(httpCode);
          Serial.print("status : ");
          Serial.println(stat);
          if (stat == "on") {
            state = SEND_DATA;
          }
        } else {
          state = CONNECT;
        }

        http.end(); //Free the resources
      }

      break;
    case SEND_DATA:
      if (millis() >= time_now + period) {
        time_now += period;
        state = QUESTION;
        if ((WiFi.status() == WL_CONNECTED)) { //Check the current connection status

          HTTPClient http;

          http.begin("http://192.168.137.1:8080/data"); //Specify the URL
          http.addHeader("Content-Type", "text/plain"); //Specify content-type header

          int minuteInt = minute();
          String minuteStr = String(minuteInt);
          if (minuteInt < 10)
            minuteStr = "0" + minuteStr;


          int secondInt = second();
          String secondStr = String(secondInt);
          if (secondInt < 10)
            secondStr = "0" + secondStr;

          String timeNow = String(day()) + "-" + String(month()) + "-" + String(year())  + " " + String(hour()) + ":" + minuteStr + ":" + secondStr;
          String message = "{\"id\":\"" + id + "\",\"sensor\":\"" + sensorName + "\",\"value\": \"" + sensorVal + "\",\"time\":\"" + timeNow + "\"}";
          String encoded = jwt.encodeJWT(message);
          int httpCode = http.POST(encoded); //Make the request

          if (httpCode > 0) { //Check for the returning code

            String payload = http.getString();
            String decoded = "";
            jwt.decodeJWT(payload, decoded);
            Serial.println("RESPONSE : ");
            //            Serial.println(httpCode);
            Serial.println(decoded);

            count = count + 1;
            File fileToAppend = SPIFFS.open("/test.txt", FILE_APPEND);
            if (fileToAppend.println(String(count) + ". value : " + String(sensorVal) + ", " + "time : " + timeNow)) {
              Serial.println("File was written");
            } else {
              Serial.println("File write failed");
            }
            fileToAppend.close();

            state = QUESTION;
          } else {
            state = CONNECT;
          }

          http.end(); //Free the resources
        }
      }

      break;
    case RESET:
      if ((WiFi.status() == WL_CONNECTED)) { //Check the current connection status

        HTTPClient http;

        http.begin("http://192.168.137.1:8080/reset/" + id); //Specify the URL

        int httpCode = http.GET(); //Make the request

        if (httpCode > 0) { //Check for the returning code

          String payload = http.getString();

          String decoded = "";
          jwt.decodeJWT(payload, decoded);

          Serial.println("RESPONSE : ");
          //          Serial.println(httpCode);
          Serial.println(decoded);

          ssidFlag = true;
          passFlag = false;

          SPIFFS.remove("/test.txt");

          state = CONNECT;
        } else {
          state = CONNECT;
        }

        http.end(); //Free the resources
      }

      break;
    case SYNC:
      if ((WiFi.status() == WL_CONNECTED)) { //Check the current connection status

        HTTPClient http;

        http.begin("http://192.168.137.1:8080/sync/" + id); //Specify the URL
        int httpCode = http.GET();                                        //Make the request

        if (httpCode > 0) { //Check for the returning code

          String payload = http.getString();

          String decoded = "";
          jwt.decodeJWT(payload, decoded);

          const size_t capacitySync = JSON_OBJECT_SIZE(2) + 40;
          DynamicJsonBuffer jsonBuffer(capacitySync);
          JsonObject& root = jsonBuffer.parseObject(decoded);

          String timeSync = root["time"].as<String>(); // "21-04-2020 18:47:27"
          int total = root["total"]; // 3

          char dateTimeSync[18] = {};
          timeSync.toCharArray(dateTimeSync, 18);
          createElements(dateTimeSync);
          setTime(makeTime(tm));//set Ardino system clock to compiled time

          Serial.println("RESPONSE : ");
          //          Serial.println(httpCode);
          Serial.print("time : ");
          Serial.println(timeSync);
          Serial.print("total : ");
          Serial.println(total);

          count = total;
          state = QUESTION;
        } else {
          state = CONNECT;
        }

        http.end(); //Free the resources
      }

      break;
    case HISTORY:

      if (hisFlag) {
        Serial.println("Please input n!!!");
      } else {
        if ((WiFi.status() == WL_CONNECTED)) { //Check the current connection status

          HTTPClient http;

          http.begin("http://192.168.137.1:8080/history/" + id + "/" + n); //Specify the URL
          int httpCode = http.GET();                                        //Make the request

          if (httpCode > 0) { //Check for the returning code

            String payload = http.getString();

            String decoded = "";
            jwt.decodeJWT(payload, decoded);

            const size_t capacityHis = JSON_OBJECT_SIZE(3) + 60;
            DynamicJsonBuffer jsonBuffer(capacityHis);
            JsonObject& rootHis = jsonBuffer.parseObject(decoded);

            String sensor = rootHis["sensor"]; // "Heat"
            String val = rootHis["value"]; // "23.5"
            String timeHis = rootHis["time"]; // "01-02-2006 15:04:05"

            Serial.println("RESPONSE : ");
            //          Serial.println(httpCode);
            Serial.print("sensor : ");
            Serial.println(sensor);
            Serial.print("value : ");
            Serial.println(val);
            Serial.print("time : ");
            Serial.println(timeHis);
            state = QUESTION;
          } else {
            state = CONNECT;
          }

          http.end(); //Free the resources
        }
      }
      break;
  }
}


void createElements(const char *str)
{
  sscanf(str, "%d-%d-%d %d:%d:%d", &Day, &Month, &Year, &Hour, &Minute, &Second);
  tm.Year = CalendarYrToTm(Year);
  tm.Month = Month;
  tm.Day = Day;
  tm.Hour = Hour;
  tm.Minute = Minute;
  tm.Second = Second;
}
