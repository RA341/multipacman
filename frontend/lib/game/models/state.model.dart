import 'dart:convert';

class GameStateModel {
  final Set<(double, double)> ghostsEaten;
  final Set<(double, double)> pelletsEaten;
  final Set<(double, double)> powerUpsEaten;
  final String type;
  final String playerSecretToken;
  final String controllingSpriteId;

  GameStateModel({
    required this.controllingSpriteId,
    required this.playerSecretToken,
    required this.ghostsEaten,
    required this.pelletsEaten,
    required this.powerUpsEaten,
    required this.type,
  });

  factory GameStateModel.fromRawJson(String str) =>
      GameStateModel.fromJson(json.decode(str));

  String toRawJson() => json.encode(toJson());

  factory GameStateModel.fromJson(Map<String, dynamic> json) => GameStateModel(
        controllingSpriteId: json['spriteId'] as String,
        playerSecretToken: json["secretToken"] as String,
        ghostsEaten: Set.from(json["ghostsEaten"].map((x) => (x[0], x[1]))),
        pelletsEaten: Set.from(json["pelletsEaten"].map((x) => (x[0], x[1]))),
        powerUpsEaten: Set.from(
          json["powerUpsEaten"].map((x) => (x[0], x[1])),
        ),
        type: json["type"],
      );

  Map<String, dynamic> toJson() => {
        "ghostsEaten": List<dynamic>.from(ghostsEaten.map((x) => x)),
        "pelletsEaten": List<dynamic>.from(pelletsEaten.map((x) => x)),
        "powerUpsEaten": List<dynamic>.from(powerUpsEaten.map((x) => x)),
        "type": type,
        playerSecretToken: playerSecretToken
      };
}
