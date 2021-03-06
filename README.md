# IITK - COIN
The aim of this project‌ is ‌to‌ ‌build‌ ‌a‌ ‌**pseudo-coin‌ ‌system‌** ‌(centralized)‌ ‌for‌ ‌use‌ ‌in‌ ‌the‌‌ IITK‌ ‌Campus.‌ To fulfill this purpose, at first back-end system of the project is being developed.

## Table of Contents

1. [General info](#general-info)
2. [Technologies](#technologies)
4. [Status](#status)

## General Info

1. The‌ ‌currency‌ ‌is‌ ‌imagined‌ ‌to‌‌ be‌‌ similar‌‌ to‌‌ premium‌‌ currencies‌‌ in‌‌ video‌‌ games‌‌ (for ‌‌example,‌‌Primogems ‌‌in‌‌ Genshin ‌‌Impact ) ‌‌with‌‌ the ‌‌additional‌‌ feature‌ ‌that‌ ‌it‌ ‌can‌ ‌be‌ ‌traded‌ ‌between‌ ‌people‌ ‌while‌ ‌paying‌ ‌a‌ ‌small‌‌ amount‌ ‌of‌ ‌tax.‌
2. Unlike Bitcoin, the amount that can be generated is not limited. As long as SnT conducts events, there will be avenues to earn. Moreover, the currency is not de-centralised and will be under control of a central authority (GenSec and Associate Heads).
3. This currency will be regulated by the Council Core Team.

## Technologies

1. **Programming language:** GO
    * GO ‌makes‌ ‌an‌ ‌ideal‌ ‌choice‌ ‌for‌ ‌backend‌ ‌web‌ ‌development,‌‌ particularly‌ ‌for‌ high-performing‌ ‌concurrent‌ ‌services‌ ‌on‌ ‌the‌‌ server‌ ‌sides.
    * Golang is fast and easy to learn.
    * One of the best feature of Golang is its ability to support concurrency. The Go language has Goroutines, which are basically functions that can run simultaneously and independently.
2. **Testing tool:** Postman
    * Since‌ ‌the‌ ‌initial goal‌ ‌of‌ ‌this ‌project‌ ‌is‌ ‌to‌ ‌build‌ ‌a‌ ‌backend,‌ and thus there‌ ‌is‌‌ no‌ ‌frontend‌ ‌UI‌ ‌for‌ ‌now, so I have used the robust framework of Postman to test the APIs.
3. **Database:** SQLite
    * SQLite is a very light weighted database so, it is easy to use it as an embedded software.
    * Reading and writing operations are very fast for SQLite database. It is almost 35% faster than File system.
    * SQLite is very easy to learn. We don't need to install and configure it. Just download SQLite libraries in the computer and it is ready for creating the database.
    * It can be used with all programming languages without any compatibility issue.
4. **Shipping:** Docker
    * Docker is a set of platform as a service products that use OS-level virtualization to deliver software in packages called containers. Containers are isolated from one another and bundle their own software, libraries and configuration files. They can communicate with each other through well-defined channels.


## Status

1. A basic `\signup` endpoint has been programmed which takes a new user input ("rollno", "name", "batch", "IsAdmin", and "password").
2. Then the password is hashed and salted and is stored in a table `Auth` containing "rollno" as primary key and "hashed_password". Rest all user details is stored in `User` table with initial coin set to 0.
3. There is a `\login` endpoint which takes user's rollno and password, and match it with the password stored in database. If the password matches a jwt is produced having a secret key and expiration time, and get stored in cookies.
4. If the jwt present in the cookie is valid, then the user can access `\secret_page`, which further allows user to access `\award`, `\transfer`, and `\balance` endpoints.
5. The `\award` is used to add coins in the database of the user, `\transfer` is used to tansfer coins from one user to another, and `\balance` is used to get the details of how much coin the user has.
6. There are three levels of users:
    * Super Admin (GenSec's account), where IsAdmin variable stores '2'.
    * Admin (Club's account), where IsAdmin variable stores '1',
    * Rest, where IsAdmin variable stores '0'.
7. Only Super Admin is allowed to award coins to any user. Personal account of GenSec, Associate Heads, and Coordinators of different clubs and socities will be freezed.
8. There is 33% tax on inter batch transfer, and 2% tax on intra batch transfer. No tax will be imposed on any transfer from admin, and super admin.
9. Users in "Rest" category, must participate in atleast '1' event to get/transfer coins.
10. GenSec and AHs can post the items that can be redeemed, which will get stored in a table RedeemItem.
11. Students can post redeem request, which will get stored in a table RedeemPending. Once a request get processed, the corresponding row will be deleted. For each request posted, item left will decrease by 1.
12. Only GenSec and AHs can accept or reject requests. if the request is accepted, coins will get deducted from the student wallet, and if rejected corresponding value of item left will increase by 1.
13. Built a docker file for this application, created a [dockerhub repo](https://hub.docker.com/r/roushanp/iitk-coin), and connected that with this github repo.