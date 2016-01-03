# Todo

- Accept null values in JSON and CSV data and queries. (Parse the input and
  replace interface{nil} with math.NaN(). NaN values should be just ignored
  from any further calculation and mark missing data.)

- Implement batch training algorithm.

- Offer some kind of auto training mode that infers all settings.

- Implement weight mask.

- Add labels?
