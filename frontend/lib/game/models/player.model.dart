import 'dart:convert';

class PlayerModel {
  final String type;
  final String playerid;
  final String user;
  final String spriteType;
  final int x;
  final int y;
  final String extraInfo;

  PlayerModel({
    required this.type,
    required this.playerid,
    required this.user,
    required this.spriteType,
    required this.x,
    required this.y,
    required this.extraInfo,
  });

  factory PlayerModel.fromRawJson(String str) =>
      PlayerModel.fromJson(json.decode(str));

  String toRawJson() => json.encode(toJson());

  factory PlayerModel.fromJson(Map<String, dynamic> json) => PlayerModel(
        type: json["type"],
        playerid: json["playerid"],
        user: json["user"],
        spriteType: json["spriteType"],
        x: json["x"],
        y: json["y"],
        extraInfo: json["extraInfo"],
      );

  Map<String, dynamic> toJson() => {
        "type": type,
        "playerid": playerid,
        "user": user,
        "spriteType": spriteType,
        "x": x,
        "y": y,
        "extraInfo": extraInfo,
      };
}
