# Roadmap

## TODO Bootstrap realistic guesses

After doing a calibration with test data, use the `ProbabilitiesOfBestGuess` and get a minimum value of the final probability (informedness * probability) that is a decent best guess. In my experiments so far, using mean(ProbabilitiesOfBestGuess) - 2*std(ProbabilitiesOfBestGuess) got about 65%, which might considered a good minimum, but it could be dependent on a lot of things so should be calculated for each family.