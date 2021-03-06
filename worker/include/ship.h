/*
** Title:ship.h
**
** Description:
**  Ship record
**
** Development Environment:
**   Ubuntu 18.04.3 LTS (Bionic Beaver)
**   gcc version 7.5.0 (Ubuntu 7.5.0-3ubuntu1-18.04)
**
** Author:
**   G.S. Cole (guycole at gmail dot com)
*/
#include <string>

#include "player.h"
#include "utility.h"

#ifndef SHIP_H_
#define SHIP_H_

class Ship {
    bool active;
    std::string id; // UUID
    Player owner;

    int warp_engines;
    int impulse_engines;
    int photon_torpedo_tubes;
    int phaser_banks;
    int deflector_shields;
    int computer;
    int life_support;
    int radio;
    int tractor_beam;

    static const std::string kBlueScouts[];
    static const std::string kBlueFighters[];
    static const std::string kBlueMiners[];
    static const std::string kBlueFlagships[];

    static const std::string kRedScouts[];
    static const std::string kRedFighters[];
    static const std::string kRedMiners[];
    static const std::string kRedFlagships[];

    void get_ship_name(char *results, PlayerTeam team, ShipType ship_type);

    public:
        Ship();

        bool is_active() {return active;}
        void set_active() {active = true;}
        void set_inactive() {active = false;}

        std::string get_id() {return id;}
        void set_id(std::string arg) {id = arg;}

//        PlayerTeam get_team() {return team;}
//        void set_team(PlayerTeam arg) {team = arg;}

        void fresh_ship(std::string id, PlayerTeam team);

        void dump_player();
};

#endif