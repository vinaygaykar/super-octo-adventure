package main;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.nio.file.Path;

public class CubeConundrum {

    public record CubeSet(int red, int green, int blue) {
    }

    public record Game(int id, CubeSet[] sets) {
    }

    public static void main(final String... args) throws IOException {
        if (args.length < 1)
            System.err.println("invalid number of arguments passed, pass in the game information file path");

        final var path = Paths.get(args[0]);
        final var games = parseAllGames(path);

        var sum = 0;
        var sumOfPower = 0;
        for (final var game : games) {
            if (isGamePossible(game, 12, 13, 14)) {
                sum += game.id;
            }

            final var set = lowestPossiblCubeSet(game);
            sumOfPower += (set.red * set.green * set.blue);
        }

        System.out.println(sum);
        System.out.println(sumOfPower);
    }

    private static Game[] parseAllGames(final Path path) throws IOException {
        return Files.lines(path)
                .map(CubeConundrum::parseGame)
                .toArray(Game[]::new);
    }

    private static Game parseGame(final String gameStr) {
        final var gameInfo = gameStr.split(": ");
        final var id = Integer.parseInt(gameInfo[0].substring(5));

        final var setsInfo = gameInfo[1].split("; ");
        final var sets = new CubeSet[setsInfo.length];
        var ptr = 0;
        for (final var set : setsInfo) {
            var red = 0;
            var green = 0;
            var blue = 0;

            for (final var cubeStr : set.split(", ")) {
                final var cubeInfo = cubeStr.split(" ");
                final var count = Integer.parseInt(cubeInfo[0]);
                switch (cubeInfo[1]) {
                    case "red":
                        red = count;
                        break;
                    case "blue":
                        blue = count;
                        break;
                    default:
                    case "green":
                        green = count;
                        break;
                }
            }

            sets[ptr++] = new CubeSet(red, green, blue);
        }

        return new Game(id, sets);
    }

    private static boolean isGamePossible(final Game game, final int maxRed, final int maxGreen, final int maxBlue) {
        for (final var set : game.sets) {
            if (maxRed < set.red || maxGreen < set.green || maxBlue < set.blue) {
                return false;
            }
        }

        return true;
    }

    private static CubeSet lowestPossiblCubeSet(final Game game) {
        var red = 0;
        var green = 0;
        var blue = 0;

        for (final var set : game.sets) {
            red = Math.max(red, set.red);
            green = Math.max(green, set.green);
            blue = Math.max(blue, set.blue);
        }

        return new CubeSet(red, green, blue);
    }

}
