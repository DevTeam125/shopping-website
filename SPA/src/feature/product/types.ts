export type ProductFull = {
  id: string;
  name: string;
  price: Price;
  status: Status;
  rating: Rating;
  description: string;
  photos: Array<{
    title: string;
    url: string;
  }>;
  features: Array<{
    title: string;
    description: string;
  }>;
};

export type Product = {
  id: string;
  name: string;
  price: Price;
  status: Status;
  rating: Rating;
  photo: {
    title: string;
    url: string;
  };
};

type Price = number;
export type Rating = 1 | 2 | 3 | 4 | 5;
type Status = 'available' | 'pending' | 'sold';
