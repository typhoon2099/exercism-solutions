//
// This is only a SKELETON file for the 'Twelve Days' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

const DAYS = {
  first: 'and a Partridge in a Pear Tree',
  second: 'two Turtle Doves',
  third: 'three French Hens',
  fourth: 'four Calling Birds',
  fifth: 'five Gold Rings',
  sixth: 'six Geese-a-Laying',
  seventh: 'seven Swans-a-Swimming',
  eighth: 'eight Maids-a-Milking',
  ninth: 'nine Ladies Dancing',
  tenth: 'ten Lords-a-Leaping',
  eleventh: 'eleven Pipers Piping',
  twelfth: 'twelve Drummers Drumming',
}

export const recite = (start, end) => {
  return Object.keys(DAYS).slice(start - 1, end ?? start).map(day => {
    return `On the ${day} day of Christmas my true love gave to me: ${giftsForDay(day)}.\n`;
  }).join("\n")
};

const giftsForDay = (day) => {
  if (day === 'first') {
    return DAYS[day].substring(4);
  }
  
  let gifts = Object.values(DAYS).slice(0, Object.keys(DAYS).indexOf(day) + 1);

  return gifts.reverse().join(', ');
}