<template>
    <nav class="bottom-nav">
      <div
        v-for="item in navItems"
        :key="item.label"
        class="nav-item"
        :class="{ active: currentRoute === item.route }"
        @click="navigate(item.route)"
      >
        <!-- If using inline SVG (optional) -->
        <svg
          v-if="item.inlineSvg"
          class="icon"
          viewBox="0 0 24 24"
          fill="currentColor"
        >
          <g v-html="item.inlineSvg" />
        </svg>
  
        <!-- If using a local SVG file -->
        <img
          v-else-if="item.svgSrc"
          class="icon"
          :src="item.svgSrc"
          :alt="item.label + ' icon'"
        />
  
        <span class="label">{{ item.label }}</span>
      </div>
    </nav>
  </template>
  
  <script lang="ts">
   import { defineComponent } from "vue";
   import type { PropType } from "vue";
   export default defineComponent({
    name: 'BottomNav',
    props: {
      navItems: {
        type: Array,
        default: () => []
      },
      currentRoute: {
        type: String,
        default: ''
      }
    },
    methods: {
      navigate(route:any) {
        // Replace this with your preferred routing logic.
        // If using Vue Router, for example:
        if (this.$router) {
          this.$router.push(route);
        } else {
          // Fallback or custom logic
          window.location.href = route;
        }
      }
    }
  });
  </script>
  
  <style scoped>
.bottom-nav {
  display: flex;
  justify-content: space-around;
  align-items: center;
  height: 64px;
  width: 100%;
  position: fixed;
  bottom: 0;
  left: 0;
  background-color: #fff;
  box-shadow: 0 -2px 8px rgba(0, 0, 0, 0.05);
  z-index: 999;
}

.nav-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  flex: 1;
  padding: 6px 0;
  color: #888;
  cursor: pointer;
  transition: color 0.2s;
}

.nav-item.active {
  color: #2f80ed; /* or your brand’s primary color */
}

.nav-item:hover {
  color: #2f80ed;
}

.icon {
  width: 24px;
  height: 24px;
  margin-bottom: 4px;
}

.label {
  font-size: 0.75rem;
}
.nav-item:hover img.icon,
.nav-item.active img.icon {
  /* Adjust the filter values to match your desired color (e.g., #2f80ed) */
  filter: brightness(0) saturate(100%) invert(39%) sepia(75%) saturate(2323%) hue-rotate(205deg) brightness(101%) contrast(102%);
}
  </style>
  