<template>
  <div class="icon-grid">
    <div v-for="(item, index) in items" :key="index" class="icon-item">
      <div class="icon-circle">
        <img :src="item.iconSrc" :alt="item.label" class="icon-img" />
      </div>
      <p class="icon-label">{{ item.label }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import recieptIcon from "../assets/reciept.svg"
import surveyIcon from "../assets/survey.svg"
import qrScanner from "../assets/qrscanner.svg"
import buyPoints from "../assets/buy.svg"
import inviteFriends from "../assets/invite.svg"
import cashPoints from "../assets/payment.svg"

const items = ref([
  { label: "Receipt Photo", iconSrc: recieptIcon },
  { label: "Survey", iconSrc: surveyIcon },
  { label: "Scan QR Code", iconSrc: qrScanner },
  { label: "Buy Points", iconSrc: buyPoints },
  { label: "Invite Friends", iconSrc:inviteFriends },
  { label: "Cashing Points", iconSrc: cashPoints },
]);
</script>

<style scoped>
/* Grid Layout */
.icon-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(120px, 1fr));
  gap: 20px;
  justify-items: center;
  margin: auto;
  padding: 20px;
  margin-bottom: 20px;
  border: 1px solid var(--border-color, #e0e0e0);
  border-radius: 8px;
  background-color: var(--card-background, #fff);
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.05);
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.icon-grid:hover {
  transform: translateY(-4px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

/* Individual Icon Item */
.icon-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  transition: transform 0.2s ease-in-out;
}

.icon-item:hover {
  transform: scale(1.05);
}

/* Circular Icon Background */
.icon-circle {
  width: 80px;
  height: 80px;
  background: #fff;
  border: 1px solid #e0e0e0;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.12);
  transition: box-shadow 0.3s ease;
}

.icon-circle:hover {
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.2);
}

/* Icon Image */
.icon-img {
  width: 40px;
  height: 40px;
  object-fit: contain;
}

/* Label Below Icons */
.icon-label {
  margin-top: 10px;
  font-size: 14px;
  font-weight: 600;
  color: #333;
}

/* Responsive Design */
@media (max-width: 480px) {
  .icon-grid {
    grid-template-columns: repeat(2, 1fr);
  }
  
  .icon-circle {
    width: 70px;
    height: 70px;
  }
  
  .icon-img {
    width: 35px;
    height: 35px;
  }
  
  .icon-label {
    font-size: 12px;
  }
}

@media (max-width: 768px) {
  .icon-grid {
    grid-template-columns: repeat(3, 1fr);
  }
}
</style>
