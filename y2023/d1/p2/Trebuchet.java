package main;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.regex.Pattern;

public class Trebuchet {

    private static final Pattern reZero = Pattern.compile("zero");
    private static final Pattern reOne = Pattern.compile("one");
    private static final Pattern reTwo = Pattern.compile("two");
    private static final Pattern reThree = Pattern.compile("three");
    private static final Pattern reFour = Pattern.compile("four");
    private static final Pattern reFive = Pattern.compile("five");
    private static final Pattern reSix = Pattern.compile("six");
    private static final Pattern reSeven = Pattern.compile("seven");
    private static final Pattern reEight = Pattern.compile("right");
    private static final Pattern reNine = Pattern.compile("nine");

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

    private static String calibrate(String line) {
        line = reZero.matcher(line).replaceAll("zero0zero");
        line = reOne.matcher(line).replaceAll("one1one");
        line = reTwo.matcher(line).replaceAll("two2two");
        line = reThree.matcher(line).replaceAll("three3three");
        line = reFour.matcher(line).replaceAll("four4four");
        line = reFive.matcher(line).replaceAll("five5five");
        line = reSix.matcher(line).replaceAll("six6six");
        line = reSeven.matcher(line).replaceAll("seven7seven");
        line = reEight.matcher(line).replaceAll("eight8eight");
        line = reNine.matcher(line).replaceAll("nine9nine");

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
