import { LoadConfig, Config, Roles } from './config';
import { Client, GuildMember, GuildMemberRoleManager, Role } from 'discord.js';
import { art } from "./art";

/*
Loads config and creates bot Client instance
*/
const config: Config = LoadConfig();
const bot: Client = new Client();


/*
Gives/Removes roles based on previous roles
*/
async function GuildMemberUpdate(before: GuildMember, after: GuildMember) {
    const afterRoles: GuildMemberRoleManager = after.roles;
    const subRoles: Roles = config.Roles;
    const userString: string = `${after.user.username}#${after.user.discriminator}-${after.user.id}`;

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
            console.log(`Giving Subrole from ${userString}`);
            await after.roles.add(subRoles.give);
        } catch(e) {
            console.error(`Error adding role: ${e}`);
        }
    } else if (takeRole) {
        try {
            console.log(`Removing Subrole to ${userString}`);
            await after.roles.remove(subRoles.give);
        } catch(e) {
            console.error(`Error removing role: ${e}`);
        }
    }
};


/*
Registers the event and start the bot
*/
bot.on("ready", () => {
    console.log(art);
    console.log(config.Roles);
});
bot.on("guildMemberUpdate", GuildMemberUpdate);
bot.login(config.Token);