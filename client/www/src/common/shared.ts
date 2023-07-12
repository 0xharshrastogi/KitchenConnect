export interface Address {
  street: string;
  city: string;
  state: string;
  zipCode: string;
  CountryCode: string;
}

export interface User {
  firstName: string;
  lastName: string;
  email: string;
  address: Address;
}
