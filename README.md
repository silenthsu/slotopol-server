
# Slot server

[![GitHub release](https://img.shields.io/github/v/release/slotopol/server.svg)](https://github.com/slotopol/server/releases/latest)
[![Hits-of-Code](https://hitsofcode.com/github/slotopol/server?branch=main)](https://hitsofcode.com/github/slotopol/server/view?branch=main)

Slots games server. Provides functionality for Megajack, Novomatic, BetSoft, and some others slot games.

# How to build

1. Install [Golang](https://go.dev/dl/) of last version.
2. Clone project and download dependencies.
3. Build project with script at `task` directory.

For Windows command prompt:

```cmd
git clone https://github.com/slotopol/server.git
cd server
go mod download && go mod verify
task\build-win-x64.cmd
```

or for Linux shell or git bash:

```sh
git clone https://github.com/slotopol/server.git
cd server
go mod download && go mod verify
sudo chmod +x ./task/*.sh
./task/build-linux-x64.sh
```

Then web-service can be started:

```cmd
slot_win_x64 web
```

# Architecture and logic

**Database.** Service reads common database tables on start and store to database only changes. Service instance oriented on monopoly usage of it's database.

**Clubs.** There is can be served several clubs. Each club have its own undepended bank, jackpot fund with rate to this fund from spins, and deposit. Bank of club is current balance of club to which arrives coins from users spins, and from which they gets a wins. There is have linkage of users wins to bank: if bank have not enough coins to pay the win during users spins, this win combination will be skipped. Deposit of club does not used in games, it can be useful to transfer the coins from bank to fix the yield.

**Users accounts.** Accounts have registrations data only. Each account can be associated with several clubs. Each user can have individual balance to gamble for each club, and individual access rights.

# How to use HTTP API

Any API endpoints can receive data in JSON, XML, YAML, or TOML format, depended by `Content-Type` header. If `Content-Type` header not given, JSON will be used to decode as default. `Accept` header if it given, defined response data format. If it absent, will be used same format as at request.

In most cases used `POST`-method of HTTP.

Any error response have HTTP status >= 400 and object at body contains `what` field with message and unique source point error `code`.

## Without authorization

First of all you can get a list of games supported by server. This call can be without authorization.

```sh
curl -H "Content-Type: application/json" -X GET localhost:8080/gamelist
```

The response can be followed:

```json
["attila","bananasgobahamas","bananasplash","beetlemania","beetlemaniadeluxe","champagne","chicago","columbus","columbusdeluxe","dolphinspearl","dolphinspearldeluxe","dynastyofming","gryphonsgold","hottarget","jewels","jokerdolphin","justjewels","katana","kingofcards","luckyladyscharm","luckyladyscharmdeluxe","marcopolo","pharaonsgold2","pharaonsgold3","plentyontwenty","polarfox","royaltreasures","secretforest","sizzlinghot","sizzlinghotdeluxe","slotopol","slotopoldeluxe","themoneygame","unicornmagic"]
```

`/ping`, `/servinfo` and `/memusage`, `/signup` and `/signin` endpoints also does not expects authorization.

## Authorization and join to game

* Sign-in, and use token from response with any followed calls.

```sh
curl -H "Content-Type: application/json" -d '{"email":"player@example.org","secret":"Et7oAm"}' -X POST localhost:8080/signin
```

You can use token `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ3ODE4OTYyNTUsIm9yaWdfaWF0IjoxNzA2MDU2MjU1LCJ1aWQiOjN9.zpOQjWZxWrNB2tDfIc8JloX30jgO3HM9jkRXFrPjwZY` for test purpose, it given for user with UID=3 on 100 years.

* Join to game. GID received at response will be used at all calls for access to this game instance. Also you will get initial screen, and user balance at this club.

```sh
curl -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ3ODE4OTYyNTUsIm9yaWdfaWF0IjoxNzA2MDU2MjU1LCJ1aWQiOjN9.zpOQjWZxWrNB2tDfIc8JloX30jgO3HM9jkRXFrPjwZY" -d '{"cid":1,"uid":3,"alias":"jokerdolphin"}' -X POST localhost:8080/game/join
```

* Change selected bet lines. Argument `sbl` is a bitset with selected lines, 1st bit in bitset means 1st line. So, value `62` sets lines 1, 2, 3, 4, 5.

```sh
curl -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ3ODE4OTYyNTUsIm9yaWdfaWF0IjoxNzA2MDU2MjU1LCJ1aWQiOjN9.zpOQjWZxWrNB2tDfIc8JloX30jgO3HM9jkRXFrPjwZY" -d '{"gid":1,"sbl":62}' -X POST localhost:8080/game/sbl/set
```

* Make a spin. Spin returns `sid` - spin ID, by this ID it can be found at the log; `screen` with new symbols after spin; `wins` with list of win on each line if it was; `fs` - free spins remained; `gain` - total gain after spin, that can be gambled on double up; `wallet` - user balance after spin with won coins.

```sh
curl -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ3ODE4OTYyNTUsIm9yaWdfaWF0IjoxNzA2MDU2MjU1LCJ1aWQiOjN9.zpOQjWZxWrNB2tDfIc8JloX30jgO3HM9jkRXFrPjwZY" -d '{"gid":1}' -X POST localhost:8080/game/spin
```

* Double-up. If presents `gain` after spin, it can be multiplied by gamble. `mult` at argument is multiplier, and it will be `2` for red-black cards game. Returned `gain` will be multiplied on win, and zero on lose. `wallet` represents new balance of user.

```sh
curl -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ3ODE4OTYyNTUsIm9yaWdfaWF0IjoxNzA2MDU2MjU1LCJ1aWQiOjN9.zpOQjWZxWrNB2tDfIc8JloX30jgO3HM9jkRXFrPjwZY" -d '{"gid":1,"mult":2}' -X POST localhost:8080/game/doubleup
```

* Collect the gain. After win on spin, or after double-up gain can be collected. In most cases it will be collected automatically on new spin.

```sh
curl -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ3ODE4OTYyNTUsIm9yaWdfaWF0IjoxNzA2MDU2MjU1LCJ1aWQiOjN9.zpOQjWZxWrNB2tDfIc8JloX30jgO3HM9jkRXFrPjwZY" -d '{"gid":1}' -X POST localhost:8080/game/collect
```

* Get current state of game.

```sh
curl -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ3ODE4OTYyNTUsIm9yaWdfaWF0IjoxNzA2MDU2MjU1LCJ1aWQiOjN9.zpOQjWZxWrNB2tDfIc8JloX30jgO3HM9jkRXFrPjwZY" -d '{"gid":1}' -X POST localhost:8080/game/state
```

## Work with user account

Any calls for some user account can be done by another user with admin access level.

* Register new user. E-mail and secret key (password) are expected, name can be omitted. Receives user ID on success.

```sh
curl -H "Content-Type: application/json" -d '{"email":"rob@example.org","secret":"jpTyD4","name":"rob"}' -X POST localhost:8080/signup
```

* Rename user.

```sh
curl -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ3ODE4OTYyNTUsIm9yaWdfaWF0IjoxNzA2MDU2MjU1LCJ1aWQiOjN9.zpOQjWZxWrNB2tDfIc8JloX30jgO3HM9jkRXFrPjwZY" -d '{"uid":3,"name":"erigine"}' -X POST localhost:8080/user/rename
```

* Change secret key.

```sh
curl -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ3ODE4OTYyNTUsIm9yaWdfaWF0IjoxNzA2MDU2MjU1LCJ1aWQiOjN9.zpOQjWZxWrNB2tDfIc8JloX30jgO3HM9jkRXFrPjwZY" -d '{"uid":3,"oldsecret":"Et7oAm","newsecret":"pGjKsd"}' -X POST localhost:8080/user/secret
```

* Delete user. Delete-call removes account from database, move all remained user's coin to deposit, and removes all users games from database.

```sh
curl -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ3ODE4OTYyNTUsIm9yaWdfaWF0IjoxNzA2MDU2MjU1LCJ1aWQiOjN9.zpOQjWZxWrNB2tDfIc8JloX30jgO3HM9jkRXFrPjwZY" -d '{"uid":3,"secret":"Et7oAm"}' -X POST localhost:8080/user/delete
```

---
(c) schwarzlichtbezirk, 2024.
