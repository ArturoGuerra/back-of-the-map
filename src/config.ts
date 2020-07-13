import * as yaml from "js-yaml";
import * as fs from "fs";

export interface Roles {
    check: string[]
    give:  string
}

export interface Config {
    Token: string
    Roles: Roles
    Guild: string
}


function loadRoles(fpath: string): Roles {
    console.log(`Loading roles from: ${fpath}`);
    let raw = yaml.safeLoad(fs.readFileSync(fpath, 'utf8'));
    if (!raw.roles) process.exit(1);

    const roles: Roles = raw.roles;

    return roles;
}

export function LoadConfig(): Config {
    let token: string = process.env.TOKEN || process.exit(1);
    let guild: string = process.env.GUILD_ID || process.exit(1);
    let cfgPath: string = process.env.CFG_PATH || "./config.yaml"

    let roles: Roles = loadRoles(cfgPath);

    let c: Config = {
       Token: token,
       Roles: roles,
       Guild: guild,
    }

    return c;
}