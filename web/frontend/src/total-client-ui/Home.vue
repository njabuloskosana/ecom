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
      <div class="tabs-container">
        <!-- Tabs -->
        <div class="tabs">
          <button
            class="tab-button"
            :class="{ active: activeTab === 'request' }"
            @click="activeTab = 'request'"
          >
            Request
          </button>
          <button
            class="tab-button"
            :class="{ active: activeTab === 'active' }"
            @click="activeTab = 'active'"
          >
            Active
          </button>
          <button
            class="tab-button"
            :class="{ active: activeTab === 'expired' }"
            @click="activeTab = 'expired'"
          >
            Expired/Used
          </button>
        </div>
  
        <!-- Tab Content -->
        <div class="tab-content">
          <!-- Request Form -->
          <div v-if="activeTab === 'request'" class="tab-pane">
            <h2>Request Gas</h2>
            <form @submit.prevent="submitForm" class="request-form">
              <div class="form-group">
                <label for="slipNumber">Slip Number</label>
                <input
                  type="text"
                  id="slipNumber"
                  v-model="slipNumber"
                  placeholder="Enter Slip Number"
                  required
                />
              </div>
  
              <div class="form-group">
                <label for="gasKg">Gas in KGs</label>
                <input
                  type="number"
                  id="gasKg"
                  v-model="gasKg"
                  placeholder="Enter Gas Quantity"
                  min="0.1"
                  step="0.1"
                  required
                />
              </div>
  
              <button type="submit" class="submit-button">Submit Request</button>
            </form>
          </div>
  
          <!-- Active Tab -->
          <div v-if="activeTab === 'active'" class="tab-pane">
            <p v-if="activeVouchers.length === 0">No active items available.</p>
            <div v-else>
              <div class="voucher-card" v-for="voucher in activeVouchers" :key="voucher.id">
                <div class="voucher-info">
                  <img src="../assets/ewallet.gif" alt="Voucher Image" class="voucher-image" />
                  <p class="voucher-title">{{ voucher.voucherName }}</p>
                  <p class="voucher-description">Code: {{ voucher.voucherCode }}</p>
                  <p class="voucher-description" v-if="voucher.discountType === 'POINTS'">
                    Points: {{ voucher.remainingPoints }}
                  </p>
                  <p class="voucher-description" v-if="voucher.discountType === 'PERCENTAGE'">
                    Discount: {{ voucher.assignedPercentageAmount }}%
                  </p>
                  <p class="voucher-description" v-if="voucher.discountType === 'MONETARY'">
                    Value: R{{ voucher.assignedMonetaryAmount }}
                  </p>
                  <p class="expiry-date">
                    Expires on: {{ new Date(voucher.expiryDate).toLocaleDateString() }}
                  </p>
                </div>
                <!-- QR Code: clicking it opens an enlarged view -->
                <div class="qr-code-container">
                  <img
                    :src="`https://api.qrserver.com/v1/create-qr-code/?data=${voucher.voucherCode}&size=100x100`"
                    alt="Voucher QR Code"
                    class="qr-code"
                    @click="openQRModal(voucher.voucherCode)"
                  />
                </div>
              </div>
            </div>
          </div>
  
          <!-- Expired/Used Tab -->
          <div v-if="activeTab === 'expired'" class="tab-pane">
            <p v-if="expiredVouchers.length === 0">No expired/used items available.</p>
            <div v-else>
              <div class="voucher-card" v-for="voucher in expiredVouchers" :key="voucher.id">
                <div class="voucher-info">
                  <img src="../assets/alert_with_coin.gif" alt="Voucher Image" class="voucher-image" />
                  <p class="voucher-title">{{ voucher.voucherName }}</p>
                  <p class="voucher-description">Code: {{ voucher.voucherCode }}</p>
                  <p class="voucher-description" v-if="voucher.discountType === 'POINTS'">
                    Points: {{ voucher.remainingPoints }}
                  </p>
                  <p class="voucher-description" v-if="voucher.discountType === 'PERCENTAGE'">
                    Discount: {{ voucher.assignedPercentageAmount }}%
                  </p>
                  <p class="voucher-description" v-if="voucher.discountType === 'MONETARY'">
                    Value: R{{ voucher.assignedMonetaryAmount }}
                  </p>
                  <p class="expiry-date">
                    Expired on: {{ new Date(voucher.expiryDate).toLocaleDateString() }}
                  </p>
                </div>
                <span
                  class="voucher-status"
                  :class="{
                    redeemed: voucher.claimed,
                    expired: !voucher.isActive && !voucher.claimed
                  }"
                >
                  {{ voucher.claimed ? 'Redeemed' : 'Expired' }}
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
  
      <footer class="footer">
        <div class="user-info">
          <!-- <div class="points">
              <img src="../assets/coins.png" alt="coins" class="coins" />
              {{ `+ 0 points` }}
            </div> -->
          <div class="points">
            <img src="../assets/gas.png" alt="coins" class="coins" />
            {{ `0 Kgs` }}
          </div>
          <button class="logout-button" @click="logout">
            <img src="../assets/logout.png" alt="Logout" class="logout-icon" />
            Logout
          </button>
        </div>
      </footer>
  
      <!-- QR Code Modal for enlarged view -->
      <div v-if="isQRModalVisible" class="qr-modal" @click="closeQRModal">
        <div class="qr-modal-content" @click.stop>
          <img :src="currentQRCode" alt="QR Code Enlarged" class="qr-modal-image" />
          <button class="qr-modal-close" @click="closeQRModal">Close</button>
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref, onMounted } from "vue";
  import { useRouter } from "vue-router";
  import NotificationModal from "../components/Notification.vue";
  import type { Voucher } from "@/interfaces";
  import Config from "@/config";
  const baseUrl = Config.BASE_API_URL_PRODUCTION;
  
  const activeTab = ref("request");
  const slipNumber = ref("");
  const gasKg = ref("");
  
  // Reactive state variables
  const isErrorModalVisible = ref<boolean>(false);
  const errorMessage = ref<string>("You are about to exit the portal. Are you sure?");
  const activeVouchers = ref<Array<Voucher>>([
    /*{
      id: 1,
      uuId: "abc123-xyz789",
      userId: 101,
      voucherId: 5001,
      voucherName: "Holiday Discount",
      voucherCode: "HOLIDAY50",
      usersFullName: "John Doe",
      discountType: "Percentage",
      claimed: false,
      remainingPoints: 100,
      assignedPercentageAmount: 50,
      assignedMonetaryAmount: 0,
      isActive: true,
      expiryDate: "2025-12-31",
    },
    {
      id: 2,
      uuId: "def456-uvw123",
      userId: 102,
      voucherId: 5002,
      voucherName: "Loyalty Bonus",
      voucherCode: "LOYALTY100",
      usersFullName: "Jane Smith",
      discountType: "Monetary",
      claimed: true,
      remainingPoints: 0,
      assignedPercentageAmount: 0,
      assignedMonetaryAmount: 100,
      isActive: false,
      expiryDate: "2024-06-30",
    },
    {
      id: 3,
      uuId: "ghi789-mno456",
      userId: 103,
      voucherId: 5003,
      voucherName: "New Year Special",
      voucherCode: "NY2025",
      usersFullName: "Michael Johnson",
      discountType: "Percentage",
      claimed: false,
      remainingPoints: 75,
      assignedPercentageAmount: 25,
      assignedMonetaryAmount: 0,
      isActive: true,
      expiryDate: "2025-01-31",
    },*/
  ]);
  
  const expiredVouchers = ref<Array<Voucher>>([
   /* {
      id: 4,
      uuId: "jkl012-pqr345",
      userId: 104,
      voucherId: 5004,
      voucherName: "Summer Sale",
      voucherCode: "SUMMER20",
      usersFullName: "Alice Brown",
      discountType: "Monetary",
      claimed: true,
      remainingPoints: 0,
      assignedPercentageAmount: 0,
      assignedMonetaryAmount: 20,
      isActive: false,
      expiryDate: "2023-08-31",
    },
    {
      id: 5,
      uuId: "mno345-stu678",
      userId: 105,
      voucherId: 5005,
      voucherName: "Winter Discount",
      voucherCode: "WINTER15",
      usersFullName: "Bob White",
      discountType: "Percentage",
      claimed: false,
      remainingPoints: 0,
      assignedPercentageAmount: 15,
      assignedMonetaryAmount: 0,
      isActive: false,
      expiryDate: "2023-12-31",
    },
    {
      id: 6,
      uuId: "pqr678-vwx901",
      userId: 106,
      voucherId: 5006,
      voucherName: "Spring Offer",
      voucherCode: "SPRING10",
      usersFullName: "Charlie Green",
      discountType: "Points",
      claimed: true,
      remainingPoints: 50,
      assignedPercentageAmount: 0,
      assignedMonetaryAmount: 0,
      isActive: false,
      expiryDate: "2023-09-30",
    },*/
  ]);
  
  const router = useRouter();
  
  // Validate authentication on component mount
  onMounted(() => {
    validateAuth();
    const userInfo = localStorage.getItem("userInfo");
    if (userInfo) {
    const parsedUser = JSON.parse(userInfo);
    const userId = parsedUser.uuId;
    getAllUserAssignedVouchers(userId);
    }
  });
  
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
  
  const submitForm = () => {
    alert(`Slip Number: ${slipNumber.value}\nGas in KGs: ${gasKg.value}`);
    // Reset form fields
    slipNumber.value = "";
    gasKg.value = "";
  };
  
  /* --- New QR Modal functionality --- */
  const isQRModalVisible = ref(false);
  const currentQRCode = ref("");
  const openQRModal = (code: string) => {
    // Set a larger QR code size for scanning
    currentQRCode.value = `https://api.qrserver.com/v1/create-qr-code/?data=${code}&size=300x300`;
    isQRModalVisible.value = true;
  };
  const closeQRModal = () => {
    isQRModalVisible.value = false;
  };

  const getAllUserAssignedVouchers = async (userId: string) => {
  try {
    const response = await fetch(`${baseUrl}/Voucher/Users/${userId}`);
    if (!response.ok) {
      throw new Error("Network response was not ok");
    }
    const res = await response.json();
    if (res.length > 0) {
      activeVouchers.value = res.filter(
        (voucher: Voucher) => voucher.isActive && !voucher.claimed
      );
      expiredVouchers.value = res.filter(
        (voucher: Voucher) => !voucher.isActive || voucher.claimed
      );
    } else {
      activeVouchers.value = [];
      expiredVouchers.value = [];
    }
 

  } catch (error) {
    console.error("Error fetching vouchers", error);
  }
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
  
  /* Updated: Make the tabs container scrollable with space for the fixed footer */
  .tabs-container {
    width: 100%;
    max-width: 600px;
    margin: auto;
    margin-top: 20px;
    font-family: Arial, sans-serif;
    color: #333;
    flex: 1;
    overflow-y: auto;
    padding-bottom: 80px; /* Reserve space for the footer */
  }
  
  /* Footer now fixed at the bottom */
  .footer {
    display: flex;
    align-items: center;
    padding: 10px 20px;
    background: #333;
    color: white;
    border-top: 1px solid #ddd;
    position: fixed;
    bottom: 0;
    left: 0;
    width: 100%;
    z-index: 10;
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
    width: 20px;
    height: 35px;
  }
  
  .points {
    display: flex;
    align-items: center;
    gap: 5px;
    font-size: 16px;
    font-weight: bold;
    color: white;
  }
  
  /* Tabs */
  .tabs {
    display: flex;
    border-bottom: 2px solid #ddd;
  }
  
  .tab-button {
    flex: 1;
    padding: 12px 20px;
    border: none;
    background: none;
    font-size: 16px;
    cursor: pointer;
    transition: 0.3s ease-in-out;
    border-bottom: 3px solid transparent;
  }
  
  .tab-button:hover {
    background: #f5f5f5;
  }
  
  .tab-button.active {
    font-weight: bold;
    color: #007bff;
    border-bottom: 3px solid #007bff;
  }
  
  /* Tab Content */
  .tab-content {
    padding: 20px;
    text-align: center;
  }
  
  .tab-pane {
    animation: fadeIn 0.3s ease-in-out;
  }
  
  /* Request Form */
  .request-form {
    display: flex;
    flex-direction: column;
    gap: 15px;
    text-align: left;
    max-width: 400px;
    margin: auto;
  }
  
  .form-group {
    display: flex;
    flex-direction: column;
  }
  
  label {
    font-weight: bold;
    margin-bottom: 5px;
  }
  
  input {
    padding: 10px;
    font-size: 16px;
    border: 1px solid #ccc;
    border-radius: 5px;
    transition: 0.2s;
  }
  
  input:focus {
    border-color: #007bff;
    outline: none;
    box-shadow: 0 0 5px rgba(0, 123, 255, 0.3);
  }
  
  .submit-button {
    padding: 12px;
    font-size: 16px;
    color: white;
    background: #007bff;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    transition: 0.3s;
  }
  
  .submit-button:hover {
    background: #0056b3;
  }
  
  /* Fade Animation */
  @keyframes fadeIn {
    from {
      opacity: 0;
      transform: translateY(5px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }
  
  .voucher-card {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 16px;
    margin-bottom: 12px;
    border-radius: 10px;
    background-color: #ffffff;
    border: 1px solid #e0e0e0;
    color: #333333;
    width: 100%;
    box-sizing: border-box;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  }
  
  .voucher-info {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
  }
  
  .voucher-image {
    width: 50px;
    height: 50px;
    border-radius: 50%;
    margin-bottom: 8px;
  }
  
  .voucher-title {
    font-weight: 600;
    font-size: 16px;
    margin: 0;
    color: #333333;
  }
  
  .voucher-description {
    font-size: 14px;
    color: #666666;
    margin: 0;
  }
  
  .expiry-date {
    font-size: 12px;
    color: #999999;
    margin-top: 4px;
  }
  
  .voucher-status {
    font-size: 14px;
    padding: 4px 8px;
    border-radius: 16px;
  }
  
  .voucher-status.redeemed {
    background-color: #1976d2;
    color: white;
  }
  
  .voucher-status.expired {
    background-color: #e57373;
    color: white;
  }
  
  .qr-code-container {
    margin-left: 20px;
    cursor: pointer;
  }
  
  .qr-code {
    width: 100px;
    height: 100px;
    border: 1px solid #ddd;
    border-radius: 8px;
  }
  
  /* QR Modal Styles */
  .qr-modal {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.7);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 20;
  }
  
  .qr-modal-content {
    background: white;
    padding: 20px;
    border-radius: 8px;
    text-align: center;
    position: relative;
  }
  
  .qr-modal-image {
    width: 300px;
    height: 300px;
  }
  
  .qr-modal-close {
    margin-top: 10px;
    padding: 8px 16px;
    background: #007bff;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
  }
  </style>
  
  