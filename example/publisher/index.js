const express = require("express");
const cors = require("cors");
const axios = require("axios");

const app = express();
const PORT = process.env.PORT || 3000;

axios.defaults.baseURL = "http://localhost:3003";

const service = axios.create({
    timeout: 30000,
});

app.use(
    cors({
        origin: "*",
    })
);

app.use(express.json({ limit: "25mb" }));

app.route("/publish").get(async (req, res) => {
    await service({
        url: `/publish?topic=topic1`,
        method: "GET",
        data: {
            message: "Hello World!",
        }
    })

    res.send("Hello World!");
});

app.listen(PORT, () => console.log("Server Started!"));