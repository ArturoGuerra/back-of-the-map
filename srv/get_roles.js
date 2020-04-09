const discord = require("discord.js");

const client = new discord.Client();

client.on("ready", () => {
    let guild = client.guilds.cache.get(process.env.GUILD_ID)
    if (!guild) process.exit(1);

    guild.roles.cache.forEach(async (item) => {
        console.log("------------------");
        console.log(item.id);
        console.log(item.name);
        console.log("------------------");
    });
});

client.login(process.env.TOKEN);