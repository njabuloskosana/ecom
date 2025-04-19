<template>
    <div class="bike-listings-container">
      <div class="header">
        <h1>Bike Listings</h1>
        <p>Find available delivery bikes or manage your fleet</p>
        
        <div class="tabs">
          <div 
            class="tab" 
            :class="{ active: activeTab === 'marketplace' }" 
            @click="activeTab = 'marketplace'"
          >
            Marketplace
          </div>
          <div 
            class="tab" 
            :class="{ active: activeTab === 'my-fleet' }" 
            @click="activeTab = 'my-fleet'"
          >
            My Fleet
          </div>
        </div>
        
        <div class="add-bike-button">
          <button @click="addNewBike">
            <span class="filter-icon">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <polygon points="22 3 2 3 10 12.46 10 19 14 21 14 12.46 22 3"/>
              </svg>
            </span>
            Add New Bike
          </button>
        </div>
      </div>
  
      <!-- Marketplace Tab Content -->
      <div v-if="activeTab === 'marketplace'" class="marketplace-content">
        <div class="marketplace-header">
          <h2>Find the Perfect Delivery Bike</h2>
          <p>Filter available bikes by your requirements</p>
        </div>
  
        <div class="filter-section">
          <div class="filter-group">
            <label>Location</label>
            <div class="filter-dropdown" @click="toggleLocationDropdown">
              <div class="dropdown-header">
                <span class="location-icon">
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0 1 18 0z"/>
                    <circle cx="12" cy="10" r="3"/>
                  </svg>
                </span>
                <span>{{ selectedLocation }}</span>
                <span class="dropdown-arrow">▼</span>
              </div>
              <div class="location-dropdown-menu" v-if="locationDropdownOpen">
                <div 
                  class="dropdown-item"
                  v-for="location in locations" 
                  :key="location"
                  :class="{ selected: location === selectedLocation }"
                  @click="selectLocation(location)"
                >
                  <span v-if="location === selectedLocation" class="check-icon">✓</span>
                  {{ location }}
                </div>
              </div>
            </div>
          </div>
  
          <div class="filter-group">
            <label>Price Range</label>
            <div class="filter-dropdown">
              <div class="dropdown-header">
                <span>{{ selectedPriceRange }}</span>
                <span class="dropdown-arrow">▼</span>
              </div>
            </div>
          </div>
  
          <div class="filter-group">
            <label>Bike Type</label>
            <div class="filter-dropdown">
              <div class="dropdown-header">
                <span>{{ selectedBikeType }}</span>
                <span class="dropdown-arrow">▼</span>
              </div>
            </div>
          </div>
        </div>
  
        <div class="bike-grid">
          <div v-for="bike in availableBikes" :key="bike.id" class="bike-card">
            <div class="bike-image">
              <span>Bike Image</span>
            </div>
            <div class="bike-info">
              <h3>{{ bike.name }}</h3>
              <div class="availability-tag" :class="{ available: bike.status === 'Available' }">
                {{ bike.status }}
              </div>
              <div class="bike-location">
                <span class="location-icon">
                  <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0 1 18 0z"/>
                    <circle cx="12" cy="10" r="3"/>
                  </svg>
                </span>
                <span>{{ bike.location }}</span>
              </div>
              <div class="bike-price">R{{ bike.price }}/month</div>
              <div class="bike-features">
                <div v-for="feature in bike.features" :key="feature" class="feature">
                  <span class="check-icon">✓</span>
                  <span>{{ feature }}</span>
                </div>
              </div>
              <div class="bike-provider">
                Listed by: {{ bike.provider }}
              </div>
              <button class="view-details-button" @click="viewBikeDetails(bike.id)">
                View Details
              </button>
            </div>
          </div>
        </div>
      </div>
      
      <!-- My Fleet Tab Content -->
      <div v-if="activeTab === 'my-fleet'" class="fleet-content">
        <div class="fleet-header">
          <h2>Manage Your Fleet</h2>
          <p>Track and maintain your bikes for optimal performance</p>
        </div>
  
        <div class="fleet-table">
          <div class="table-header">
            <div class="bike-details">Bike Details</div>
            <div class="location">Location</div>
            <div class="status">Status</div>
            <div class="assigned-to">Assigned To</div>
            <div class="maintenance">Maintenance</div>
            <div class="actions">Actions</div>
          </div>
  
          <div class="table-body">
            <div v-for="bike in fleetBikes" :key="bike.id" class="table-row">
              <div class="bike-details">
                <div class="bike-model">{{ bike.model }}</div>
                <div class="bike-id">ID #{{ bike.id }}</div>
              </div>
              <div class="location">{{ bike.location }}</div>
              <div class="status">
                <div class="status-tag" :class="{ rented: bike.status === 'Rented', available: bike.status === 'Available' }">
                  {{ bike.status }}
                </div>
              </div>
              <div class="assigned-to">{{ bike.assignedTo || '—' }}</div>
              <div class="maintenance">
                <div class="maintenance-status" :class="bike.maintenanceStatus.toLowerCase().replace(' ', '-')">
                  {{ bike.maintenanceStatus }}
                </div>
              </div>
              <div class="actions">
                <button class="edit-button" @click="editBike(bike.id)">Edit</button>
                <button class="view-button" @click="viewBike(bike.id)">View</button>
              </div>
            </div>
          </div>
        </div>
  
        <div class="fleet-actions">
          <button class="download-report-button" @click="downloadFleetReport">
            Download Fleet Report
          </button>
          <button class="schedule-maintenance-button" @click="scheduleMaintenanceModal = true">
            Schedule Maintenance
          </button>
        </div>
  
        <div class="performance-metrics">
          <h2>Performance Metrics</h2>
          <p>Track your fleet's efficiency and earnings</p>
  
          <div class="metrics-grid">
            <div class="metric-card">
              <h3>Fleet Utilization</h3>
              <div class="metric-value">67%</div>
              <div class="metric-detail">2 of 3 bikes rented</div>
            </div>
  
            <div class="metric-card">
              <h3>Monthly Revenue</h3>
              <div class="metric-value">R1350</div>
              <div class="metric-detail">From 2 active rentals</div>
            </div>
  
            <div class="metric-card">
              <h3>Maintenance Costs</h3>
              <div class="metric-value">R450</div>
              <div class="metric-detail">Last 30 days</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent, ref } from 'vue';
  
  interface Bike {
    id: string;
    name: string;
    model: string;
    status: string;
    location: string;
    price: number;
    features: string[];
    provider: string;
    assignedTo?: string;
    maintenanceStatus: string;
  }
  
  export default defineComponent({
    name: 'BikeListings',
    setup() {
      const activeTab = ref('my-fleet');
      
      // Locations for the dropdown
      const locations = ref(['All Locations', 'Johannesburg', 'Pretoria', 'Centurion', 'Midrand']);
      const selectedLocation = ref('All Locations');
      const locationDropdownOpen = ref(false);
      
      const selectedPriceRange = ref('All Prices');
      const selectedBikeType = ref('All Types');
  
      // Available bikes for marketplace
      const availableBikes = ref<Bike[]>([
        {
          id: '101',
          name: 'Honda Delivery 125cc',
          model: 'Honda Delivery 125cc',
          status: 'Available',
          location: 'Johannesburg',
          price: 650,
          features: ['Low fuel consumption', 'Cargo box included', 'Recently serviced'],
          provider: 'SpeedFleet Delivery',
          maintenanceStatus: 'Good condition'
        },
        {
          id: '102',
          name: 'Yamaha YBR 125',
          model: 'Yamaha YBR 125',
          status: 'Available',
          location: 'Pretoria',
          price: 700,
          features: ['New tires', 'Phone mount included', 'GPS tracking'],
          provider: 'Urban Couriers',
          maintenanceStatus: 'Due for service'
        },
        {
          id: '103',
          name: 'Suzuki GS 150',
          model: 'Suzuki GS 150',
          status: 'Available',
          location: 'Centurion',
          price: 600,
          features: ['Large delivery box', 'Helmet included', 'Rain gear provided'],
          provider: 'FastTrack Deliveries',
          maintenanceStatus: 'Recently serviced'
        },
        {
          id: '104',
          name: 'TVS Star 110',
          model: 'TVS Star 110',
          status: 'Available',
          location: 'Midrand',
          price: 600,
          features: ['Economical', 'Lightweight', 'Easy to handle'],
          provider: 'EcoDelivery Solutions',
          maintenanceStatus: 'Good condition'
        }
      ]);
  
      // Fleet bikes for my fleet
      const fleetBikes = ref<Bike[]>([
        {
          id: '101',
          name: 'Honda Delivery 125cc',
          model: 'Honda Delivery 125cc',
          status: 'Rented',
          location: 'Johannesburg',
          price: 650,
          features: ['Low fuel consumption', 'Cargo box included', 'Recently serviced'],
          provider: 'SpeedFleet Delivery',
          assignedTo: 'Michael Brown',
          maintenanceStatus: 'Good condition'
        },
        {
          id: '102',
          name: 'Yamaha YBR 125',
          model: 'Yamaha YBR 125',
          status: 'Rented',
          location: 'Pretoria',
          price: 700,
          features: ['New tires', 'Phone mount included', 'GPS tracking'],
          provider: 'Urban Couriers',
          assignedTo: 'Lisa Smith',
          maintenanceStatus: 'Due for service'
        },
        {
          id: '103',
          name: 'Suzuki GS 150',
          model: 'Suzuki GS 150',
          status: 'Available',
          location: 'Centurion',
          price: 600,
          features: ['Large delivery box', 'Helmet included', 'Rain gear provided'],
          provider: 'FastTrack Deliveries',
          maintenanceStatus: 'Recently serviced'
        }
      ]);
  
      const scheduleMaintenanceModal = ref(false);
  
      function toggleLocationDropdown() {
        locationDropdownOpen.value = !locationDropdownOpen.value;
      }
  
      function selectLocation(location: string) {
        selectedLocation.value = location;
        locationDropdownOpen.value = false;
      }
  
      function addNewBike() {
        console.log('Add new bike clicked');
        // Implementation for adding a new bike
      }
  
      function viewBikeDetails(bikeId: string) {
        console.log('View bike details clicked for bike ID:', bikeId);
        // Implementation for viewing bike details
      }
  
      function editBike(bikeId: string) {
        console.log('Edit bike clicked for bike ID:', bikeId);
        // Implementation for editing a bike
      }
  
      function viewBike(bikeId: string) {
        console.log('View bike clicked for bike ID:', bikeId);
        // Implementation for viewing a bike
      }
  
      function downloadFleetReport() {
        console.log('Download fleet report clicked');
        // Implementation for downloading fleet report
      }
  
      return {
        activeTab,
        locations,
        selectedLocation,
        locationDropdownOpen,
        selectedPriceRange,
        selectedBikeType,
        availableBikes,
        fleetBikes,
        scheduleMaintenanceModal,
        toggleLocationDropdown,
        selectLocation,
        addNewBike,
        viewBikeDetails,
        editBike,
        viewBike,
        downloadFleetReport
      };
    }
  });
  </script>
  
  <style scoped>
  .bike-listings-container {
    font-family: Arial, sans-serif;
    padding: 20px;
    max-width: 1200px;
    margin: 0 auto;
    color: #333;
  }
  
  .header {
    display: flex;
    flex-direction: column;
    margin-bottom: 20px;
    position: relative;
  }
  
  .header h1 {
    font-size: 24px;
    margin: 0 0 5px 0;
  }
  
  .header p {
    font-size: 14px;
    margin: 0 0 15px 0;
    color: #666;
  }
  
  .tabs {
    display: flex;
    border-bottom: 1px solid #ddd;
    margin-bottom: 20px;
  }
  
  .tab {
    padding: 10px 20px;
    cursor: pointer;
    border-bottom: 2px solid transparent;
  }
  
  .tab.active {
    border-bottom: 2px solid #3366ff;
    color: #3366ff;
    font-weight: bold;
  }
  
  .add-bike-button {
    position: absolute;
    right: 0;
    top: 10px;
  }
  
  .add-bike-button button {
    background-color: #1a1a2e;
    color: white;
    border: none;
    padding: 8px 16px;
    border-radius: 4px;
    cursor: pointer;
    display: flex;
    align-items: center;
    gap: 8px;
  }
  
  .filter-icon {
    display: inline-flex;
  }
  
  /* Marketplace Tab Styles */
  .marketplace-header {
    margin-bottom: 20px;
  }
  
  .marketplace-header h2 {
    font-size: 18px;
    margin: 0 0 5px 0;
  }
  
  .marketplace-header p {
    font-size: 14px;
    margin: 0;
    color: #666;
  }
  
  .filter-section {
    display: flex;
    gap: 20px;
    margin-bottom: 30px;
  }
  
  .filter-group {
    flex: 1;
  }
  
  .filter-group label {
    display: block;
    margin-bottom: 8px;
    font-size: 14px;
  }
  
  .filter-dropdown {
    border: 1px solid #ddd;
    border-radius: 4px;
    position: relative;
  }
  
  .dropdown-header {
    padding: 10px;
    display: flex;
    align-items: center;
    cursor: pointer;
  }
  
  .location-icon {
    margin-right: 8px;
    display: inline-flex;
  }
  
  .dropdown-arrow {
    margin-left: auto;
    font-size: 12px;
  }
  
  .location-dropdown-menu {
    position: absolute;
    top: 100%;
    left: 0;
    right: 0;
    border: 1px solid #ddd;
    border-top: none;
    border-radius: 0 0 4px 4px;
    background-color: white;
    z-index: 10;
  }
  
  .dropdown-item {
    padding: 10px;
    cursor: pointer;
    display: flex;
    align-items: center;
  }
  
  .dropdown-item:hover {
    background-color: #f5f5f5;
  }
  
  .dropdown-item.selected {
    background-color: #f0f7ff;
  }
  
  .check-icon {
    margin-right: 8px;
    color: #3366ff;
  }
  
  .bike-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
    gap: 20px;
  }
  
  .bike-card {
    border: 1px solid #ddd;
    border-radius: 8px;
    overflow: hidden;
  }
  
  .bike-image {
    height: 150px;
    background-color: #f5f5f5;
    display: flex;
    align-items: center;
    justify-content: center;
    color: #999;
  }
  
  .bike-info {
    padding: 15px;
  }
  
  .bike-info h3 {
    margin: 0 0 8px 0;
    font-size: 16px;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  
  .availability-tag {
    font-size: 12px;
    padding: 4px 8px;
    border-radius: 4px;
    background-color: #e0e0e0;
  }
  
  .availability-tag.available {
    background-color: #e6f7e6;
    color: #2c7a2c;
  }
  
  .bike-location {
    display: flex;
    align-items: center;
    font-size: 14px;
    margin: 8px 0;
    color: #666;
  }
  
  .bike-price {
    font-weight: bold;
    margin: 8px 0;
  }
  
  .bike-features {
    margin: 10px 0;
  }
  
  .feature {
    font-size: 13px;
    margin-bottom: 5px;
    display: flex;
    align-items: center;
  }
  
  .bike-provider {
    font-size: 12px;
    color: #666;
    margin: 10px 0;
  }
  
  .view-details-button {
    width: 100%;
    padding: 8px;
    background-color: #1a1a2e;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    margin-top: 10px;
  }
  
  /* Fleet Tab Styles */
  .fleet-header {
    margin-bottom: 20px;
  }
  
  .fleet-header h2 {
    font-size: 18px;
    margin: 0 0 5px 0;
  }
  
  .fleet-header p {
    font-size: 14px;
    margin: 0;
    color: #666;
  }
  
  .fleet-table {
    border: 1px solid #ddd;
    border-radius: 8px;
    overflow: hidden;
    margin-bottom: 20px;
  }
  
  .table-header {
    display: grid;
    grid-template-columns: 2fr 1fr 1fr 1fr 1fr 1fr;
    background-color: #f5f5f5;
    padding: 15px;
    font-weight: bold;
  }
  
  .table-body {
    display: flex;
    flex-direction: column;
  }
  
  .table-row {
    display: grid;
    grid-template-columns: 2fr 1fr 1fr 1fr 1fr 1fr;
    padding: 15px;
    border-top: 1px solid #ddd;
  }
  
  .bike-model {
    font-weight: bold;
  }
  
  .bike-id {
    font-size: 12px;
    color: #666;
    margin-top: 5px;
  }
  
  .status-tag {
    display: inline-block;
    padding: 4px 8px;
    border-radius: 4px;
    font-size: 12px;
  }
  
  .status-tag.rented {
    background-color: #e7e7e7;
    color: #333;
  }
  
  .status-tag.available {
    background-color: #e6f7e6;
    color: #2c7a2c;
  }
  
  .maintenance-status {
    display: inline-block;
    padding: 4px 8px;
    border-radius: 4px;
    font-size: 12px;
  }
  
  .maintenance-status.good-condition {
    background-color: #e6f7e6;
    color: #2c7a2c;
  }
  
  .maintenance-status.due-for-service {
    background-color: #fff9e6;
    color: #997a00;
  }
  
  .maintenance-status.recently-serviced {
    background-color: #e6f0ff;
    color: #3366cc;
  }
  
  .actions {
    display: flex;
    gap: 10px;
  }
  
  .edit-button, .view-button {
    padding: 5px 10px;
    border: 1px solid #ddd;
    background-color: white;
    border-radius: 4px;
    cursor: pointer;
  }
  
  .edit-button:hover, .view-button:hover {
    background-color: #f5f5f5;
  }
  
  .fleet-actions {
    display: flex;
    gap: 10px;
    margin-bottom: 30px;
  }
  
  .download-report-button {
    padding: 8px 16px;
    background-color: white;
    border: 1px solid #ddd;
    border-radius: 4px;
    cursor: pointer;
  }
  
  .schedule-maintenance-button {
    padding: 8px 16px;
    background-color: #1a1a2e;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
  
  .performance-metrics h2 {
    font-size: 18px;
    margin: 0 0 5px 0;
  }
  
  .performance-metrics p {
    font-size: 14px;
    margin: 0 0 20px 0;
    color: #666;
  }
  
  .metrics-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: 20px;
  }
  
  .metric-card {
    padding: 20px;
    border: 1px solid #ddd;
    border-radius: 8px;
  }
  
  .metric-card h3 {
    font-size: 14px;
    margin: 0 0 10px 0;
    color: #666;
  }
  
  .metric-value {
    font-size: 28px;
    font-weight: bold;
    margin-bottom: 5px;
  }
  
  .metric-detail {
    font-size: 12px;
    color: #666;
  }
  </style>