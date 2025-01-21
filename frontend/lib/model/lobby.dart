import 'dart:convert';

class LobbyModel {
  final int id;
  final DateTime createdAt;
  final DateTime updatedAt;
  final DateTime? deletedAt;
  final String lobbyName;
  final int userId;
  final int joined;
  final String username;

  LobbyModel({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.deletedAt,
    required this.lobbyName,
    required this.userId,
    required this.joined,
    required this.username,
  });

  factory LobbyModel.fromRawJson(String str) =>
      LobbyModel.fromJson(json.decode(str));

  String toRawJson() => json.encode(toJson());

  factory LobbyModel.fromJson(Map<String, dynamic> json) => LobbyModel(
        id: json["ID"],
        createdAt: DateTime.parse(json["CreatedAt"]),
        updatedAt: DateTime.parse(json["UpdatedAt"]),
        deletedAt: json["DeletedAt"] == null
            ? null
            : DateTime.parse(json["DeletedAt"] as String),
        lobbyName: json["LobbyName"],
        userId: json["UserID"],
        joined: json["Joined"],
        username: json["Username"],
      );

  Map<String, dynamic> toJson() => {
        "ID": id,
        "CreatedAt": createdAt.toIso8601String(),
        "UpdatedAt": updatedAt.toIso8601String(),
        "DeletedAt": deletedAt?.toIso8601String(),
        "LobbyName": lobbyName,
        "UserID": userId,
        "Joined": joined,
        "Username": username,
      };
}
