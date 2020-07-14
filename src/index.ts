import { LoadConfig, Config, Roles } from './config';
import { Client, GuildMember, GuildMemberRoleManager, Guild } from 'discord.js';
import { art } from "./art";

/*
Loads config and creates bot Client instance
*/
const config: Config = LoadConfig();
const bot: Client = new Client();

/*
Assigns or removes role based on current roles
*/
async function RoleHandler(member: GuildMember) {
    const afterRoles: GuildMemberRoleManager = member.roles;
    const userString: string = `${member.user.username}#${member.user.discriminator}-${member.user.id}`;
    const subRoles: Roles = config.Roles;
    const hasGive: boolean = afterRoles.cache.some(i => i.id == subRoles.give);
    let aftercheck: boolean = subRoles.check.some(id => {
        return afterRoles.cache.some(i => id == i.id);
    });

    // Gives role if user: has one or more sub roles after && and doesn't have cat role
    let giveRole: boolean = (aftercheck && !hasGive);

    // Takes role if user doesn't have any sub roles after && has cat role
    let takeRole: boolean = (!aftercheck && hasGive);

    if (giveRole) {
        try {
            console.log(`Giving Subrole to ${userString}`);
            await member.roles.add(subRoles.give);
        } catch(e) {
            console.error(`Error adding role: ${e}`);
        }
    } else if (takeRole) {
        try {
            console.log(`Removing Subrole from ${userString}`);
            await member.roles.remove(subRoles.give);
        } catch(e) {
            console.error(`Error removing role: ${e}`);
        }
    }
}

/*
Gives/Removes roles when the bot gets the on_ready event from discord
*/
async function InitialRoles() {
    let guild: Guild = bot.guilds.cache.get(config.Guild)
    if (guild == null) process.exit(1);

    guild.members.cache.forEach(RoleHandler);
}


/*
Gives/Removes roles based on previous roles
*/
async function GuildMemberUpdate(_: GuildMember, after: GuildMember) {
    try {
       await RoleHandler(after);
    } catch (e) {
        console.error(e);
    }
};


/*
Registers the event and start the bot
*/
bot.on("ready", async () => {
    console.log(art);
    console.log(config.Roles);

    try {
        await InitialRoles();
        console.log("Done updating roles");
    } catch(e) {
        console.error("Error running initial role updater");
    }
});
bot.on("guildMemberUpdate", GuildMemberUpdate);
bot.login(config.Token);