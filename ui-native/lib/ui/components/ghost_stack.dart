import 'package:flutter/material.dart';

const baseFolder = 'assets/app_images';

const orangeGhost = '$baseFolder/ghost-orange.png';
const redGhost = '$baseFolder/ghost-red.png';
const pinkGhost = '$baseFolder/ghost-pink.png';
const blueGhost = '$baseFolder/ghost-blue.png';
const pacmanImg = '$baseFolder/pacman.png';

class GhostStack extends StatelessWidget {
  const GhostStack({super.key, required this.child});

  final Widget child;

  @override
  Widget build(BuildContext context) {
    return Stack(
      alignment: AlignmentDirectional.topCenter,
      children: [
        Positioned(
          left: 50,
          top: 100,
          child: ImageBackgroundComponent(
            faceRight: true,
            assetName: orangeGhost,
          ),
        ),
        Positioned(
          left: 80,
          top: 280,
          child: ImageBackgroundComponent(
            faceRight: true,
            assetName: pinkGhost,
          ),
        ),
        Positioned(
          right: 50,
          top: 150,
          child: ImageBackgroundComponent(
            assetName: redGhost,
          ),
        ),
        Positioned(
          left: 50,
          bottom: 10,
          child: ImageBackgroundComponent(
            assetName: pacmanImg,
            imageScale: 7,
            faceRight: true,
          ),
        ),
        Positioned(
          right: 80,
          bottom: 250,
          child: ImageBackgroundComponent(
            assetName: blueGhost,
          ),
        ),
        child,
      ],
    );
  }
}

class ImageBackgroundComponent extends StatelessWidget {
  const ImageBackgroundComponent({
    required this.assetName,
    this.faceRight = false,
    this.imageScale = 2.7,
    super.key,
  });

  final double imageScale;
  final bool faceRight;
  final String assetName;

  @override
  Widget build(BuildContext context) {
    return Transform(
      alignment: Alignment.center,
      transform: faceRight ? Matrix4.rotationY(3.1416) : Matrix4.rotationY(0),
      child: Opacity(
        opacity: 0.7,
        child: Image.asset(
          assetName,
          scale: imageScale,
        ),
      ),
    );
  }
}
