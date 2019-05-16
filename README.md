## Gosu server

this is mostly just a refactoring of my previous attempt at creating this service, as it makes use of a different file structure and packages

### TODO
- Migrate to time.Time for `History.MatchDate` and `Challenge.ResolutionDate`, Then adjust contest resolution functions to make use of the new types

### Current Routes
- `/auth/login GET` redirects to bnet oauth login screen
- `/auth/bnet_oauth_cb` parses information returned bnet and awards a jwt token if successful
- `/auth/user GET` returns player information of current user

requires admin permissions
- `/player/{id}/add-admin PUT` returns a player of a given ID with raised permissions assuming the caller has the ability to do so
- `/player/{id}/remove-admin PUT` returns a player of a given ID with demoted permissions assuming the caller has the ability to do so
- `/player/{id}/promote/{rank} PUT` returns a list of players affected by the change of promoting a player of a given ID and Rank
- `/player/{id}/replace/{rank} PUT` returns a list of players affected by the change of changing a player's rank to given rank, even if it displaces the former holder of said rank
- `/player/{id}/register PUT` returns a player of a given ID with a populated ranking for settle the beef
- `/player/{id}/unregister PUT` returns a list of players affected by the change of removing given player from the rankings

does not require admin privelidges
- `/player GET` returns a list all the players
- `/player/rankings GET` returns a list all the players involved in "Settle the Beef"
- `/player/register PUT` returns a player of a given ID with a populated ranking for settle the beef
- `/player/unregister PUT` returns a list of players affected by the change of removing player from rankings
- `/player/{id} GET` returns a player with the given ID
- `/player/{id}/challenges GET`  returns a list of challenges related to a player of a given ID (both as challenger and challenged)
- `/player/{id}/history GET` returns a list of matches related to a player of a given ID

requires admin permissions
- `/challenge/{id}/adjudicate` returns a challenge that has been mutated by admin decision about the winner of the challenge

- `/challenge GET` returns a list of all challenges
- `/challenge/unresolved GET` returns a list of all challeges that have yet to resolve
- `/challenge POST` returns a new challenge after saving it to DB
- `/challenge/{id} PUT` returns a challenge of a given ID after updating some information, limit params that can be set
- `/challenge/{id} DELETE` returns a success message after soft deleting a challenge

- `/history/{id} GET` returns record of a match given an ID

This is basically an endpoint for scheduled jobs only. Updating player profile information, match history, adjudicating challenges
- `/cron`