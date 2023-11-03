const express = require("express");
const cors = require("cors");
const axios = require("axios");

const app = express();
const PORT = process.env.PORT || 3001;

axios.defaults.baseURL = "http://localhost:3003";

const service = axios.create({
    timeout: 30000,
});

service({
    url: `/register?topic=topic1`,
    method: "GET",
    headers: {
        "X-Consumer": `http://localhost:${PORT}`,
        "X-Handler-Path": "/handler",
    }
})

app.use(
    cors({
        origin: "*",
    })
);

app.use(express.json({ limit: "25mb" }));

app.route("/handler").post(async (req, res) => {
    console.log("Processing: ");
    console.log(req.body);
});

app.listen(PORT, () => console.log("Server Started!"));