# Hayasui

Hayasui is a discord bot to search and provide you anime/manga/character/people data with interactive response message.

[Hayasui](https://en.wikipedia.org/wiki/Japanese_fleet_oiler_Hayasui)'s name is taken from a japanese fleet oiler which support and resupply fuel, ammo, and food for other ships. Also, [exists](https://kancolle.fandom.com/wiki/Hayasui) in Kantai Collection games and anime. To live up its name, this bot will 'supply' you with anime, manga, character, and people information.

## Features

- Search anime/manga/character/people.
- Get anime info.
- Get manga info.
- Get character info.
- Get people info.
- Interactive response message (pagination supported).

[Sample results](https://github.com/rl404/hayasui/blob/master/sample.md)

## Requirement

- [Discord bot](https://discordpy.readthedocs.io/en/latest/discord.html) and its token
- [Go](https://golang.org/)
- [Redis](https://redis.io/) (optional)
- [Docker](https://docker.com) + [Docker compose](https://docs.docker.com/compose/) (optional)

## Steps

1. Git clone this repo.
   ```bash
   git clone github.com/rl404/hayasui
   ```
2. Rename `sample.env` to `.env` and modify according to your configuration.

| Name                  |  Default   | Description                                |
| --------------------- | :--------: | ------------------------------------------ |
| `HYS_DISCORD_TOKEN`\* |            | Discord bot token.                         |
| `HYS_DISCORD_PREFIX`  |    `>`     | Discord bot prefix command.                |
| `HYS_CACHE_DIALECT`   | `inmemory` | Cache type (`redis`,`inmemory`,`memcache`) |
| `HYS_CACHE_ADDRESS`   |            | Cache address.                             |
| `HYS_CACHE_PASSWORD`  |            | Cache passowrd.                            |
| `HYS_CACHE_TIME`      |   `24h`    | Cache duration.                            |

3. Run.

   ```bash
   make bot

   # or using docker
   make docker
   # to stop docker
   make docker-stop
   ```

4. Invite the bot to your server.
5. Try `>help`.
6. Have fun.

## Bot Commands

### Search

```bash
>search <anime|manga|character|people> <query> [query...]

# example
>search anime naruto
>search manga one piece
>search character ichigo
>search people kana
```

### Get Anime

```bash
>anime <anime_id>

# example
>anime 1
```

### Get Manga

```bash
>manga <manga_id>

# example
>manga 1
```

### Get Character

```bash
>character <character_id>

# example
>character 1
```

### Get People

```bash
>people <people_id>

# example
>people 1
```

## License

MIT License

Copyright (c) 2021 Axel
