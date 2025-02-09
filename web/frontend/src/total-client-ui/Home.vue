<template>
    <NotificationModal
      :isVisible="isErrorModalVisible"
      title="Exit Portal"
      :message="errorMessage"
      cancelText="Cancel"
      confirmText="Yes"
      @cancel="hideErrorModal"
      @confirm="confirmLogout"
    />
  
    <div class="welcome-container">
      <main>
        <h1>Welcome to the Portal</h1>
      </main>
  
      <footer class="footer">
        <div class="user-info">
          <div class="points">
            <img src="../assets/coins.png" alt="coins" class="coins" />
            {{ `3000 points` }}
          </div>
          <button class="logout-button" @click="logout">
            <img src="../assets/logout.png" alt="Logout" class="logout-icon" />
            Logout
          </button>
        </div>
      </footer>
    </div>
  </template>
  
<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import NotificationModal from "../components/Notification.vue";

// Reactive state variables
const isErrorModalVisible = ref<boolean>(false);
const errorMessage = ref<string>("You are about to exit the portal. Are you sure?");

const router = useRouter();

// Validate authentication on component mount
onMounted(() => {
    validateAuth();
});

// Methods
const validateAuth = () => {
    const userInfo = localStorage.getItem("userInfo");
    if (!userInfo) {
        router.push("/");
    }
};



const hideErrorModal = () => {
    isErrorModalVisible.value = false;
};

const logout = () => {
    isErrorModalVisible.value = true;
};
const confirmLogout = () => {
    localStorage.removeItem("userInfo");
    localStorage.removeItem("jwtToken");
    localStorage.removeItem("userProfile");
    router.push("/");
};
</script>
  
  <style scoped>
  .welcome-container {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
    text-align: center;
    background-color: #f4f4f4;
  }
  
  main {
    flex: 1;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  
  h1 {
    font-size: 2rem;
    color: #333;
  }
  
  .footer {
    display: flex;
    align-items: center;
    padding: 10px 20px;
    background: #333;
    color: white;
    border-top: 1px solid #ddd;
  }
  
  .user-info {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 10px;
    width: 100%;
  }
  
  .logout-button {
    display: flex;
    align-items: center;
    gap: 5px;
    padding: 8px 12px;
    background-color: white;
    color: black;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    font-size: 14px;
    transition: background 0.3s ease;
  }
  
  .logout-button:hover {
    background-color: #c0392b;
  }
  
  .logout-icon {
    width: 16px;
    height: 16px;
  }
  
  .coins {
    width: 50px;
    height: 30px;
  }
  
  .points {
    display: flex;
    align-items: center;
    gap: 5px;
    font-size: 16px;
    font-weight: bold;
    color: white;
  }
  </style>
  