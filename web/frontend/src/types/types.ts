export interface BikeOwner {
    fullName: string;
    businessName?: string;
    email: string;
    phoneNumber: string;
    location: string;
    numberOfBikes: number;
    password: string;
  }
  
  export interface Driver {
    firstName: string;
    lastName: string;
    email: string;
    phoneNumber: string;
    password: string;
  }

  export interface Bike {
    id: string;
    name: string;
    model: string;
    status: 'Available' | 'Rented';
    location: string;
    price: number;
    features: string[];
    provider: string;
    assignedTo?: string;
    maintenanceStatus: 'Good condition' | 'Due for service' | 'Recently serviced';
  }
  
  export interface FleetMetrics {
    utilization: {
      percentage: number;
      rentedCount: number;
      totalCount: number;
    };
    revenue: {
      amount: number;
      activeRentals: number;
      currency: string;
    };
    maintenance: {
      amount: number;
      period: string;
      currency: string;
    };
  }