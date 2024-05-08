const UserService = require("../services/user.service");
const EventService = require("../services/event.service");
const NewsService = require("../services/news.service");
const AuthService = require("../services/auth.service");

const { assert } = require("chai");

describe("ServicesTests", () => {
  describe("UserService", () => {
    it("should return an error - The ID is not specified", async () => {
      const _id = "";

      try {
        await UserService.getUserById(_id);
        assert.fail(
          "Expected an error to be thrown, but the operation succeeded"
        );
      } catch (error) {
        assert.strictEqual(error.message, "The ID is not specified");
      }
    });
  });
});
