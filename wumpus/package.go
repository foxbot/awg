// Package wumpus provides lightweight wrappers over Discord's REST and
// Gateway APIs.
//
// Be warned that the REST wrapper is not intended to be used directly
// on Discord's REST API; but rather with a REST proxy such as
// [blaze](https://github.com/yuki-bot/blaze)
//
// This is intentional; since workers are incapable of handling
// ratelimits themselves, I outsource this behavior directly to the
// proxy
//
// Note that most of this package will be incomplete abstractions over
// Discord's API. This is intentional; I am only implementing models
// as necessary for my bots to function. Should you need a model that
// is not included here, feel free to fork out or submit a merge request
//
// This is not intended to be a fully functional Discord wrapper - see
// [discordgo](https://github.com/bwmarrin/discordgo) for that instead
package wumpus
