const telephoneValidator = (value) => /^1[3|4|5|7|8]\d{9}$/.test(value);

export default {
  telephoneValidator,
};
