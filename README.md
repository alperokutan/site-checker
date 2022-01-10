# Site Status Checker

This is a site status checker that pingin site and checking for the site is up or down.


## Endpoints

* `GET`  **/users** : Get all users
* `GET`  **/users/{id}** : Get the specific user
* `POST`  **/users** : Create a new user
* `PUT`  **/users/{id}** : Update a user
* `DELETE`  **/users/{id}** : Delete a user

<br />

* `GET`  **/user/sites/{userId}** : Get all sites of user and check status
* `GET`  **/sites/{id}** : Get the specific user
* `POST`  **/sites** : Create a new user
* `PUT`  **/sites/{id}** : Update a user
* `DELETE`  **/sites/{id}** : Delete a user

## Used Libraries

* [mux](https://github.com/gorilla/mux) - Mux
* [gorm](https://gorm.io/) - Gorm

## Will be add

* Login with jwt and middleware
* Socket for real time checking
* History of site status
* Pricing with Stripe
