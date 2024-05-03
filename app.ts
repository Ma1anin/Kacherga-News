require("dotenv").config();

import mongoose from "mongoose";
import express from "express";
import fileUpload from "express-fileupload";
import bodyParser from "body-parser";

import newsRouter from "./src/routes/news.routes";
import eventRouter from "./src/routes/event.routes";
import userRouter from "./src/routes/user.routes";
import CardService from "./src/services/card.service";

const app = express();

const urlencodedParser = bodyParser.urlencoded({
  extended: false,
});

app.use(express.static("public"));
app.use(express.json());
app.use(fileUpload());
app.use("/news", newsRouter);
app.use("/event", eventRouter);
app.use("/user", userRouter);

app.set("view engine", "hbs");
app.set("views", "public/views");

app.get("/edit-account", function (req, res) {
  res.render("edit-account.hbs");
});

app.get("/register", function (req, res) {
  res.render("register.hbs");
});

app.post("/register", urlencodedParser, async function (req, res) {
  try {
    if (!req.body) throw new Error("Request body is empty!");

    await fetch("http://localhost:3000/user", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ ...req.body, role: "user" }),
    })
      .then((response) => {
        if (!response.ok) throw new Error("Create user request was not ok");
        res.status(200).redirect("/login");
        // ИСПРАВИТЬ ЭТОТ КОСТЫЛЬ!
        return;
      })
      .catch((error) => {
        console.error(error);
      });
  } catch (err) {
    console.log(err);
  }
});

app.get("/login", function (req, res) {
  res.render("login.hbs");
});

app.get("/", function (req, res) {
  res.render("main.hbs");

  // const newsContainer = document.getElementById("news-container");
  // fetch("http://localhost:3000/news")
  //   .then((response) => response.json())
  //   .then((news) => {
  //     news.forEach((item) => {
  //       const newsCard = CardService.createNewsCard(item);
  //       newsContainer.appendChild(newsCard);
  //     });
  //   });
});

async function startApp() {
  try {
    await mongoose.connect(process.env.DB_URI);
    app.listen(process.env.PORT, () =>
      console.log(`Server started on port ${process.env.PORT}`)
    );
  } catch (err) {
    console.log(err);
  }
}

startApp();
