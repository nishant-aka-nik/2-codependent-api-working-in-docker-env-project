# 2-codependent-api-working-in-docker-env-project

2 codependent web scrapping API based on **Colly Framework** working mutually in docker environment

## How to run the project

Run the docker-compose.yml by using command which will bring up all the services together

```bash
docker compose up
```

We can access the api1 using url

```bash
http://localhost:8080/geturl
```

**Request example -**

```JSON
{
    "url" : "https://www.amazon.com/LETSCOM-Trackers-Waterproof-Activity-Compatible/dp/B08HGVZYFX/"
}
```

This **api1** will take url from request JSON and then scrape the necessary details from the url then it will send the scraped data as a payload to **api2** which then write the payload to **mongoDB** (I am using free mongoDB cluster to store the data login credentials are given below)

**MongoDB login credentials**

```
Go to - https://account.mongodb.com/account/login
Email Address - sellerappassignment@clrmail.com
Password - ^8ZMg2%a5f^7
```

After storing the data **api1** will give response as below

**Response example -**

```
{
    "message": "Successfully called and created/updated data from API2 to mongoDB ",
    "messagefromapi2": "Successfully created the record in mongoDB",
    "scrapeddata": {
        "url": "https://www.amazon.com/LETSCOM-Trackers-Waterproof-Activity-Compatible/dp/B08HGVZYFX/",
        "product": {
            "name": "LETSCOM Smart Watch, GPS Running Watch Fitness Trackers with Heart Rate Monitor Step Counter Sleep Monitor, IP68 Waterproof Digital Watch Activity Tracker Compatible with iPhone Android Phones",
            "imageURL": "https://images-na.ssl-images-amazon.com/images/I/71%2BkxwGAvJL._AC_SL1500_.jpg",
            "description": "All Day Activity with GPS Tracking: Built-in GPS lets you record the real-time route for your outdoor activities without carrying a smartphone. Choose 8 between up to 14 sports modes to track your various workout accordingly. Letscom's smart watch helps keeps track of your steps, heart rate, calories burned, distance travelled, sleep quality, and more.Smart Watch compatibility: 1.3 inches color full touch screen brings you a better interactive experience, comfortable TPU band suits for the 5.5-8.7 inches adult wrist. You can now even customize your device wallpaper via VeryfitPro app, and adjust the screen brightness and change the watch faces directly on the watch.Call And Message Notifications: This watch is compatible with most iOS 8.0 & Android 4.4 above smartphones to notify you a incoming call, text message, calendar and app notifications(including Facebook, Twitter, Youtube, WhatsApp, Emails, Instagram, and more).You can even hang up incoming phone calls straight from our smartwatch.10 Days Battery Life & IP68 Waterproof: Built with a 210mAh battery that only takes 2.5 hours to charge and can be used for up to 10 days with a standby time of 30-45 days. IP68 waterproof allows you to take showers or swim in the pool with it, also supports tracking your swimming data.A Smart Watch Full of Surprises: Letscom GPS running watch also features with many practical tools, such as weather forecasting, vibration alarm clocks, stopwatch & timer, find your phone, music controller, breath guide, female health care and more. Our smart watch is full of surprises. Best for the family and friends.",
            "price": "$45.99",
            "totalReviews": 307
        }
    }
}
```

## Request and Response Example of both APIs

**api1 Request and Response Example**

Request -

```
{
    "url" : "https://www.amazon.com/LETSCOM-Trackers-Waterproof-Activity-Compatible/dp/B08HGVZYFX/"
}
```

Response -

```
{
    "message": "Successfully called and created/updated data from API2 to mongoDB ",
    "messagefromapi2": "Successfully updated the record in mongoDB",
    "scrapeddata": {
        "url": "https://www.amazon.com/LETSCOM-Trackers-Waterproof-Activity-Compatible/dp/B08HGVZYFX/",
        "product": {
            "name": "LETSCOM Smart Watch, GPS Running Watch Fitness Trackers with Heart Rate Monitor Step Counter Sleep Monitor, IP68 Waterproof Digital Watch Activity Tracker Compatible with iPhone Android Phones",
            "imageURL": "https://images-na.ssl-images-amazon.com/images/I/71%2BkxwGAvJL._AC_SL1500_.jpg",
            "description": "All Day Activity with GPS Tracking: Built-in GPS lets you record the real-time route for your outdoor activities without carrying a smartphone. Choose 8 between up to 14 sports modes to track your various workout accordingly. Letscom's smart watch helps keeps track of your steps, heart rate, calories burned, distance travelled, sleep quality, and more.Smart Watch compatibility: 1.3 inches color full touch screen brings you a better interactive experience, comfortable TPU band suits for the 5.5-8.7 inches adult wrist. You can now even customize your device wallpaper via VeryfitPro app, and adjust the screen brightness and change the watch faces directly on the watch.Call And Message Notifications: This watch is compatible with most iOS 8.0 & Android 4.4 above smartphones to notify you a incoming call, text message, calendar and app notifications(including Facebook, Twitter, Youtube, WhatsApp, Emails, Instagram, and more).You can even hang up incoming phone calls straight from our smartwatch.10 Days Battery Life & IP68 Waterproof: Built with a 210mAh battery that only takes 2.5 hours to charge and can be used for up to 10 days with a standby time of 30-45 days. IP68 waterproof allows you to take showers or swim in the pool with it, also supports tracking your swimming data.A Smart Watch Full of Surprises: Letscom GPS running watch also features with many practical tools, such as weather forecasting, vibration alarm clocks, stopwatch & timer, find your phone, music controller, breath guide, female health care and more. Our smart watch is full of surprises. Best for the family and friends.",
            "price": "$45.99",
            "totalReviews": 307
        }
    }
}
```

**api2 Request and Response Example**

Request -

```
{
    "url": "https://www.amazon.com/dp/B01AJ0LGRQnew/",
    "product": {
        "name": "Energizer AA Batteries (20 Count), Double A Max Alkaline Battery",
        "imageURL": "https://images-na.ssl-images-amazon.com/images/I/81g5wmQHyJL._AC_SL1500_.jpg",
        "description": "20-pack of Energizer MAX alkaline AA batteriesUp to 50% longer lasting than basic alkaline in demanding devicesLeak resistant-construction protects your devices from leakage of fully used batteries for up to 2 years Bonus: It’s guaranteedNon-stop energy for your non-stop family’s must-have devices—think toys, flashlights, wireless mice, remotes, and moreHolds power up to 10 years in storage—so you’re never left powerless",
        "price": "$12.14",
        "totalReviews": 456
    }
}
```

Response -

```
{
    "message": "Successfully updated the record in mongoDB"
}
```
