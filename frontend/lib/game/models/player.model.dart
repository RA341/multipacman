import 'dart:convert';

import 'package:multipacman/game/components/utils.dart';

class PlayerModel {
  final String type;
  final String playerid;
  final String user;
  final String spriteType;
  final double x;
  final double y;
  final String dir;

  PlayerModel({
    this.dir = "right",
    required this.type,
    required this.playerid,
    required this.user,
    required this.spriteType,
    required this.x,
    required this.y,
  });

  factory PlayerModel.fromRawJson(String str) =>
      PlayerModel.fromJson(json.decode(str));

  String toRawJson() => json.encode(toJson());

  factory PlayerModel.fromJson(Map<String, dynamic> json) => PlayerModel(
        type: json["type"],
        playerid: json["playerid"],
        dir: json["dir"],
        user: json["user"],
        spriteType: json["spriteType"],
        x: double.parse(json["x"]),
        y: double.parse(json["y"]),
      );

  Map<String, dynamic> toJson() => {
        "type": type,
        "playerid": playerid,
        "user": user,
        "spriteType": spriteType,
        "x": x,
        "y": y,
      };
}
