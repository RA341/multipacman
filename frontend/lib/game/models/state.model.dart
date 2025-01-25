import 'dart:convert';
import 'dart:ffi';

class GameStateModel {
  final List<(Float, Float)> ghostsEaten;
  final List<(Float, Float)> pelletsEaten;
  final List<(Float, Float)> powerUpsEaten;
  final String type;

  GameStateModel({
    required this.ghostsEaten,
    required this.pelletsEaten,
    required this.powerUpsEaten,
    required this.type,
  });

  factory GameStateModel.fromRawJson(String str) => GameStateModel.fromJson(json.decode(str));

  String toRawJson() => json.encode(toJson());

  factory GameStateModel.fromJson(Map<String, dynamic> json) => GameStateModel(
        ghostsEaten: List.from(json["ghostsEaten"].map((x) => (x[0], x[1]))),
        pelletsEaten: List.from(json["pelletsEaten"].map((x) => (x[0], x[1]))),
        powerUpsEaten: List.from(
          json["powerUpsEaten"].map((x) => (x[0], x[1])),
        ),
        type: json["type"],
      );

  Map<String, dynamic> toJson() => {
        "ghostsEaten": List<dynamic>.from(ghostsEaten.map((x) => x)),
        "pelletsEaten": List<dynamic>.from(pelletsEaten.map((x) => x)),
        "powerUpsEaten": List<dynamic>.from(powerUpsEaten.map((x) => x)),
        "type": type,
      };
}
