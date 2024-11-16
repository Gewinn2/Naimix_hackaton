<template>
<img v-if="id === 1" src="../assets/cards/majorArcana/The_Fool.png">
<img v-else-if="id === 2" src="../assets/cards/majorArcana/The_Magician.png">
<img v-else-if="id === 3" src="../assets/cards/majorArcana/The_High_Priestess.png">
<img v-else-if="id === 4" src="../assets/cards/majorArcana/The_Empress.png">
<img v-else-if="id === 5" src="../assets/cards/majorArcana/The_Emperor.png">
<img v-else-if="id === 6" src="../assets/cards/majorArcana/The_Hierophant.png">
<img v-else-if="id === 7" src="../assets/cards/majorArcana/The_Lovers.png">
<img v-else-if="id === 8" src="../assets/cards/majorArcana/The_Chariot.png">
<img v-else-if="id === 9" src="../assets/cards/majorArcana/The_Power.png">
<img v-else-if="id === 10" src="../assets/cards/majorArcana/The_Hermit.png">
<img v-else-if="id === 11" src="../assets/cards/majorArcana/Wheel_of_Fortune.png">
<img v-else-if="id === 12" src="../assets/cards/majorArcana/Justice.png">
<img v-else-if="id === 13" src="../assets/cards/majorArcana/The_Hanged_Man.png">
<img v-else-if="id === 14" src="../assets/cards/majorArcana/Death.png">
<img v-else-if="id === 15" src="../assets/cards/majorArcana/Temperance.png">
<img v-else-if="id === 16" src="../assets/cards/majorArcana/The_Devil.png">
<img v-else-if="id === 17" src="../assets/cards/majorArcana/The_Tower.png">
<img v-else-if="id === 18" src="../assets/cards/majorArcana/The_Star.png">
<img v-else-if="id === 19" src="../assets/cards/majorArcana/The_Moon.png">
<img v-else-if="id === 20" src="../assets/cards/majorArcana/The_Sun.png">
<img v-else-if="id === 21" src="../assets/cards/majorArcana/Judgement.png">
<img v-else-if="id === 22" src="../assets/cards/majorArcana/The_World.png">
<img v-else-if="id === 23" src="../assets/cards/minorArcana/Ace of Wands.png">
<img v-else-if="id === 24" src="../assets/cards/minorArcana/Two of Wands.png">
<img v-else-if="id === 25" src="../assets/cards/minorArcana/Three of Wands.png">
<img v-else-if="id === 26" src="../assets/cards/minorArcana/Four of Wands.png">
<img v-else-if="id === 27" src="../assets/cards/minorArcana/Five of Wands.png">
<img v-else-if="id === 28" src="../assets/cards/minorArcana/Six of Wands.png">
<img v-else-if="id === 29" src="../assets/cards/minorArcana/Seven of Wands.png">
<img v-else-if="id === 30" src="../assets/cards/minorArcana/Eight of Wands.png">
<img v-else-if="id === 31" src="../assets/cards/minorArcana/Nine of Wands.png">
<img v-else-if="id === 32" src="../assets/cards/minorArcana/Ten of Wands.png">
<img v-else-if="id === 33" src="../assets/cards/minorArcana/Page of Wands.png">
<img v-else-if="id === 34" src="../assets/cards/minorArcana/Knight of Wands.png">
<img v-else-if="id === 35" src="../assets/cards/minorArcana/Queen of Wands.png">
<img v-else-if="id === 36" src="../assets/cards/minorArcana/King of Wands.png">
<img v-else-if="id === 37" src="../assets/cards/minorArcana/Ace of Cups.png">
<img v-else-if="id === 38" src="../assets/cards/minorArcana/Two of Cups.png">
<img v-else-if="id === 39" src="../assets/cards/minorArcana/Three of Cups.png">
<img v-else-if="id === 40" src="../assets/cards/minorArcana/Four of Cups.png">
<img v-else-if="id === 41" src="../assets/cards/minorArcana/Five of Cups.png">
<img v-else-if="id === 42" src="../assets/cards/minorArcana/Six of Cups.png">
<img v-else-if="id === 43" src="../assets/cards/minorArcana/Seven of Cups.png">
<img v-else-if="id === 44" src="../assets/cards/minorArcana/Eight of Cups.png">
<img v-else-if="id === 45" src="../assets/cards/minorArcana/Nine of Cups.png">
<img v-else-if="id === 46" src="../assets/cards/minorArcana/Ten of Cups.png">
<img v-else-if="id === 47" src="../assets/cards/minorArcana/Page of Cups.png">
<img v-else-if="id === 48" src="../assets/cards/minorArcana/Knight of Cups.png">
<img v-else-if="id === 49" src="../assets/cards/minorArcana/Queen of Cups.png">
<img v-else-if="id === 50" src="../assets/cards/minorArcana/King of Cups.png">
<img v-else-if="id === 51" src="../assets/cards/minorArcana/Ace of Swords.png">
<img v-else-if="id === 52" src="../assets/cards/minorArcana/Two of Swords.png">
<img v-else-if="id === 53" src="../assets/cards/minorArcana/Three of Swords.png">
<img v-else-if="id === 54" src="../assets/cards/minorArcana/Four of Swords.png">
<img v-else-if="id === 55" src="../assets/cards/minorArcana/Five of Swords.png">
<img v-else-if="id === 56" src="../assets/cards/minorArcana/Six of Swords.png">
<img v-else-if="id === 57" src="../assets/cards/minorArcana/Seven of Swords.png">
<img v-else-if="id === 58" src="../assets/cards/minorArcana/Eight of Swords.png">
<img v-else-if="id === 59" src="../assets/cards/minorArcana/Nine of Swords.png">
<img v-else-if="id === 60" src="../assets/cards/minorArcana/Ten of Swords.png">
<img v-else-if="id === 61" src="../assets/cards/minorArcana/Page of Swords.png">
<img v-else-if="id === 61" src="../assets/cards/minorArcana/Knight of Swords.png">
<img v-else-if="id === 62" src="../assets/cards/minorArcana/Queen of Swords.png">
<img v-else-if="id === 63" src="../assets/cards/minorArcana/King of Swords.png">
<img v-else-if="id === 64" src="../assets/cards/minorArcana/Ace of Pentacles.png">
<img v-else-if="id === 65" src="../assets/cards/minorArcana/Two of Pentacles.png">
<img v-else-if="id === 66" src="../assets/cards/minorArcana/Three of Pentacles.png">
<img v-else-if="id === 67" src="../assets/cards/minorArcana/Four of Pentacles.png">
<img v-else-if="id === 68" src="../assets/cards/minorArcana/Five of Pentacles.png">
<img v-else-if="id === 69" src="../assets/cards/minorArcana/Six of Pentacles.png">
<img v-else-if="id === 70" src="../assets/cards/minorArcana/Seven of Pentacles.png">
<img v-else-if="id === 71" src="../assets/cards/minorArcana/Eight of Pentacles.png">
<img v-else-if="id === 72" src="../assets/cards/minorArcana/Nine of Pentacles.png">
<img v-else-if="id === 73" src="../assets/cards/minorArcana/Ten of Pentacles.png">
<img v-else-if="id === 74" src="../assets/cards/minorArcana/Page of Pentacles.png">
<img v-else-if="id === 75" src="../assets/cards/minorArcana/Knight of Pentacles.png">
<img v-else-if="id === 76" src="../assets/cards/minorArcana/Queen of Pentacles.png">
<img v-else-if="id === 77" src="../assets/cards/minorArcana/King of Pentacles.png">
</template>
<script lang="ts">
export default{
  props:{
    id: {
      type: Number,
      required: true,
    },
  }
};
</script>