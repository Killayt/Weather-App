# Weather-App

<div class='Description'> 
    <p>
        This is an implementation of the OpenWeather API for getting weather information for any city. The data is transmitted in JSON format. Of the data there, only "temp" - temperature. If you want to get data, or get more coordinate data, edit the WeatherData structure in config/config.go and the API URL and also in config/config.go.
    </p>
</div>

<div class="How to use">
    <h3>How to use</h3>
    <p>
        To use you need to have an OpenWeather API key. Create a ".apiConfig" file. In file write: 
        <code>
            {
                "ApiKey" : "YOUR API KEY"
            }
        </code>
    </p>
    <p> 
        <ol>
            <li>Start Weather-App</li>
            <li>Go to 127.0.0.1:9000</li>
            <li>Go to 127.0.0.1:9000/weather/CitytoKnowAbout
        </ol>
    </p>
</div>

<div class="A little more information">
    <h3> Little more info </h3>
    <p>
        The server is running on port 9000. You can select a port by editing the port constant in the main.go file.
    </p>
</div>
