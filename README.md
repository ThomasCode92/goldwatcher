# Goldwatcher

Keep track of the price of gold in (almost real-time) while monitoring your current gold holdings. 📈 💰

**About this Repository**<br />
🌐 Project of the _[Building GUI Applications with Fyne and Go](https://www.udemy.com/course/building-gui-applications-with-fyne-and-go-golang/?couponCode=KEEPLEARNING)_ course.<br />
⭐ A _"Gold Price Monitor"_ application with SQL database.<br />
🔍 Making API Requests, Refreshing the UI, database interactions,...

## Start the application

To start the application, ensure you have both Fyne and Go installed on your system. You can refer to the official documentation for installation instructions:

- [Fyne Documentation](https://docs.fyne.io/) - Easily Build Native Apps that Work Everywhere
- [Go Programming Language](https://go.dev/) - Build Simple, Secure and Scalable systems

If you are using [asdf](https://asdf-vm.com/), you can also install Go via the _.tool-versions_ file. More information on this can be found [here](https://asdf-vm.com/manage/configuration.html#tool-versions).<br />
Once Go and Fyne are installed, follow these steps to start the application:

```bash
git clone https://github.com/ThomasCode92/goldwatcher
cd goldwatcher    # navigate into project folder
go run .          # start the program
```

To run the unit tests, use the command `go test -v .`.

## Data Resources

The application retrieves gold price information through some API requests to [goldprice.org](https://goldprice.org/).<br />
The price data, in JSON format, is accessed via an [API endpoint](https://data-asg.goldprice.org/dbXRates/usd) while the chart is available for download as a [.png file](https://goldprice.org/charts/gold_3d_b_k_usd_x.png).
