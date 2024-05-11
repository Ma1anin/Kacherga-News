require("dotenv").config();

const chai = require("chai");
const chaiHttp = require("chai-http");
const MongoClient = require("mongodb").MongoClient;
const mongoClient = new MongoClient(process.env.DB_URL);
const { expect } = chai;

chai.use(chaiHttp);

const login = "TESTIK";
const password = "qwerty11";
const fullName = "Freakan";

describe("HTTP Request tests", () => {
  before(async () => {
    await clearDatabase();
  });

  describe("User requests", () => {
    it("should create a user with correct data", (done) => {
      chai
        .request("http://localhost:5000")
        .post("/api/register")
        .set("Content-Type", "application/json")
        .send({
          login: login,
          password: password,
          fullName: fullName,
        })
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(200);
          expect(res.body.message).to.equals("Successful registration");
          expect(res.body.success).to.be.true;
          done();
        });
    });

    it("should login to account", (done) => {
      chai
        .request("http://localhost:5000")
        .post("/api/login")
        .set("Content-Type", "application/json")
        .send({
          login: login,
          password: password,
        })
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(200);
          expect(res.body.message).to.equals("Successful account login");
          expect(res.body.success).to.be.true;
          expect(res).to.have.cookie('userId');
          done();
        });
    });

    it("should logout of the account", (done) => {
      chai
        .request("http://localhost:5000")
        .post("/api/logout")
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(200);
          expect(res.body.message).to.equals("Successful account login");
          expect(res.body.success).to.be.true;
          done();
        });
    });
  });
});

async function clearDatabase() {
  try {
    console.log("> Очистка базы данных...");

    const db = await mongoClient.db("test");
    const collection = await db.collection("users");
    const result = await collection.drop();

    console.log(result ? "> База данных успешно очищена" : "> Безуспешно");
  } catch (err) {
    console.log(err);
  }
}
