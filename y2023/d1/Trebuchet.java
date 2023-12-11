package main;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;

public class Trebuchet {

    public static void main(final String... args) throws IOException {
        if (args.length < 1)
            System.err.println("invalid number of arguments passed, pass in the calibation document file path");
        
        final var path = Paths.get(args[0]);
        final var sum = Files.lines(path)
            .map(Trebuchet::calibrate)
            .mapToInt(Integer::valueOf)
            .sum();
        System.out.println(sum);
    }

    private static String calibrate(final String line) {
        var firstDigit = '0';
        var lastDigit = '0';

        for (int l = 0, r = line.length() - 1; l < line.length(); l++, r--) {
            final var x = line.charAt(l);
            if ('0' <= x && x <= '9') {
                lastDigit = x;
            }

            final var y = line.charAt(r);
            if ('0' <= y && y <= '9') {
                firstDigit = y;
            }
        }

        return "" + firstDigit + lastDigit;
    }

}
