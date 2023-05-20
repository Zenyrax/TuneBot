## Description
TuneBot is a simple application for automatically purchasing items from stores running on Tuneportals, a common platform for Vinyl shops. Here are some features:
  - Proxy IP support
  - Out-of-stock product monitoring
    - If it can't add a product to cart, most likely because it's out of stock, it'll keep trying
  - Discord webhooks
  - Low resource usage


## Supported sites
- https://dearbornmusic.net/
- https://waterloorecords.com/
- https://vintagevinyl.com/
- https://shop.cactusrecords.net/
- https://parkavecds.com/
- Many, many more sites run on Tuneportals
  - They tend to look similar, so you'll know when  you see one after a while
  - Their product links are all structured like https://example.net/UPC/093624869559 
 
## How to install & start using
1. Go to releases
2. Download the latest exe
3. Run it
4. To keep track of checkouts, you can go to "Manage Settings" and add a Discord webhook
    * Make sure to test it first!
    * Instructions for making a Discord webhook: https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks
5. Make tasks for the product you want to bot using the instructions below
6. If you want to run many tasks, I suggest using proxies, but I don't think they rate-limit
7. Select the "Run Tasks" options and selection the task and proxy list you want to use
    * If you don't want to use any proxies, choose a blank proxy list
8. Let it run!
    * If the item you want is out of stock, it'll keep trying to add to cart until it comes back in stock

## Setting up tasks
1. Run the exe
2. Select "Manage Tasks"
3. Open the "tasks.csv" file or make a copy of "tasks.csv" and open that
    * I like to manage CSV files with Ron's Editor
- Store
  - This is the URL of the website (Ex. vintagevinyl.com)
  - Be careful not to add an unnecessary "www"
- UPC
  - This is the UPC of the product and can be found at the end of a product link (ex. 093624869559)
  - For example, for the link https://waterloorecords.com/UPC/602445401284, the UPC would be 602445401284
- Quantity
  - This is how many of the item you want to order (ex. 1)
- Shipping Option
  - This is the shipping option you want to use 
  - Leave this blank and the bot will pick the first one available
    - This will mostly like be fine for most sites
  - To select a shipping option, go to checkout on the website you want to bot and copy the what's before the comma in the list of shipping options
    - This will mainly be useful if you want expedited shipping, in-store pickup, or just want to be safe
    - You'll need an item in your cart and your address filled out to make shipping options appear
    - Example: https://media.discordapp.net/attachments/868215243057287218/1109001432260161536/image.png
  - For more advanced users, you can directly put in the shipping id in this section to skip two parts of the checkout flow
    - This will make your checkout speed noticibly faster
    - You can find this by watching network traffic during checkout
    - The id's of the shipping options are in the response of the GET request to the /SecureCartShippingOptions.json endpoint
    - Example: https://media.discordapp.net/attachments/868215243057287218/1109007583689330768/image.png
- Price Limit
  - This the maximum price of the product (ex. 65.99)
  - To have no limit, make the limit extremely high or 0
  - I doubt this will ever matter, but maybe it'll come in handy one day
- State
  - Use state abbreviations (ex. NY for New York)
- Country 
  - Use "US" for the United States
  - I'm not sure about other countries
- Card Number
  - No spaces!
- Card Month/Year
  - Both must be two characters (ex. 04 and 25 instead of 4 and 2025)
 
## Setting up proxies
1. Run the exe
2. Select "Manage Proxies"
3. Open the "proxies.txt" file or use any blank txt file
4. Put proxies in IP:AUTH or IP:AUTH:USER:PASS format

Note: Proxies aren't required
