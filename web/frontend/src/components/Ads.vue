<template>
  <div class="carousel">
    <div class="carousel-container">
      <img
        :src="images[currentIndex]"
        alt="Advertisement"
        class="carousel-image"
      />
    </div>
    
    <!-- Slide Indicators -->
    <div class="carousel-indicators">
      <span
        v-for="(image, index) in images"
        :key="index"
        :class="['indicator', { active: index === currentIndex }]"
        @click="goToSlide(index)"
      ></span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from "vue";
import rewardAdvert from '../assets/rewards-advert.png';
import rewardAdvert2 from '../assets/rewards-advert2.png';

// List your advertisement images here. Adjust the paths as needed.
const images = ref<string[]>([
  rewardAdvert,
  rewardAdvert2
]);

const currentIndex = ref<number>(0);
let intervalId: number;

// Advances to the next slide, looping back to the first image when necessary.
const nextSlide = () => {
  currentIndex.value = (currentIndex.value + 1) % images.value.length;
};

// Directly navigates to a given slide.
const goToSlide = (index: number) => {
  currentIndex.value = index;
};

// Auto-rotate the carousel every 5 seconds.
onMounted(() => {
  intervalId = setInterval(nextSlide, 5000);
});

// Clear the interval when the component is unmounted.
onUnmounted(() => {
  clearInterval(intervalId);
});
</script>

<style scoped>
.carousel {
  position: relative;
  width: 100%;
  margin: 20px auto;
  border-radius: 8px;
  overflow: hidden;
}

.carousel-container {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 300px; /* Adjust this based on your preferred height */
  overflow: hidden;
}

.carousel-image {
  width: auto;
  max-width: 100%;
  height: 100%;
  object-fit: contain; /* Use 'cover' if you want it to fill the space */
}

/* Indicators */
.carousel-indicators {
  position: absolute;
  bottom: 10px;
  width: 100%;
  text-align: center;
}

.indicator {
  display: inline-block;
  width: 10px;
  height: 10px;
  margin: 0 4px;
  background-color: #ddd;
  border-radius: 50%;
  cursor: pointer;
  transition: background-color 0.3s;
}

.indicator.active {
  background-color: #333;
}
</style>
  