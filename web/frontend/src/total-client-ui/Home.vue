<template>
    <div class="welcome-container">
      <main>
        <h1>Welcome to the Portal</h1>
      </main>
  
      <footer>
        <button class="logout-button" @click="logout">Logout</button>
      </footer>
    </div>
  </template>
  
  <script>
export default {
    created() {
        this.validateAuth();
    },
    methods: {
        validateAuth() {
            const userInfo = localStorage.getItem("userInfo");
            if (!userInfo) {
                this.$router.push("/");
            }
        },
        logout() {
            alert("Logging out...");
            localStorage.removeItem("userInfo");
            localStorage.removeItem("jwt");
            localStorage.removeItem("profile");
            // Implement actual logout logic (e.g., clear auth tokens, redirect)
            this.$router.push("/");
        },
    },
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
  
  footer {
    background: #333;
    color: white;
    padding: 10px;
    text-align: center;
  }
  
  .logout-button {
    background: #e74c3c;
    color: white;
    border: none;
    padding: 10px 20px;
    font-size: 1rem;
    cursor: pointer;
    border-radius: 5px;
    transition: 0.3s;
  }
  
  .logout-button:hover {
    background: #c0392b;
  }
  </style>
  