/*
** Title:player_message_tell.h
**
** Description:
**
** Development Environment:
**   Ubuntu 18.04.3 LTS (Bionic Beaver)
**   gcc version 7.5.0 (Ubuntu 7.5.0-3ubuntu1-18.04)
**
** Author:
**   G.S. Cole (guycole at gmail dot com)
*/
#include "player_message.h"

#ifndef PLAYER_MESSAGE_TELL_H_
#define PLAYER_MESSAGE_TELL_H_

class PlayerMessageTell: public PlayerMessage {

    public:
        PlayerMessageTell();
};

#endif