<template>
    <div class="carousel">
      <div class="carousel-container">
        <img
          src="../assets/rewards-advert.png"
          alt="Advertisement"
          class="carousel-image"
        />
      </div>
      <!-- Navigation Buttons -->
      <button class="carousel-control prev" @click="prevSlide">&#10094;</button>
      <button class="carousel-control next" @click="nextSlide">&#10095;</button>
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
  
  // List your advertisement images here. Adjust the paths as needed.
  const images = ref<string[]>([
    "assets/rewards-advert.png"
  ]);
  
  const currentIndex = ref<number>(0);
  let intervalId: number;
  
  // Advances to the next slide, wrapping to the first image if needed.
  const nextSlide = () => {
    currentIndex.value = (currentIndex.value + 1) % images.value.length;
  };
  
  // Moves to the previous slide, wrapping to the last image if needed.
  const prevSlide = () => {
    currentIndex.value =
      (currentIndex.value - 1 + images.value.length) % images.value.length;
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
    border: 1px solid var(--border-color);
    border-radius: 8px;
    background-color: var(--card-background);
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.05);
    overflow: hidden;
  }
  
  .carousel-container {
    display: flex;
    align-items: center;
    justify-content: center;
  }
  
  .carousel-image {
    width: 100%;
    height: auto;
    transition: opacity 0.5s ease-in-out;
  }
  
  /* Navigation Buttons */
  .carousel-control {
    position: absolute;
    top: 50%;
    transform: translateY(-50%);
    background-color: rgba(255, 255, 255, 0.7);
    border: none;
    border-radius: 50%;
    width: 40px;
    height: 40px;
    cursor: pointer;
    font-size: 24px;
    line-height: 40px;
    text-align: center;
    color: #333;
    z-index: 1;
    outline: none;
    transition: background-color 0.3s;
  }
  
  .carousel-control:hover {
    background-color: rgba(255, 255, 255, 1);
  }
  
  .carousel-control.prev {
    left: 10px;
  }
  
  .carousel-control.next {
    right: 10px;
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
  
  /* Responsive Design */
  @media (max-width: 768px) {
    .carousel {
      width: 95%;
    }
  }
  </style>
  