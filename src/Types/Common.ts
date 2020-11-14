export interface IAnalize {
  dnalizeId: number;
  date: string;
  name: string;
  value: string;
  patientID: number;
}

export interface IAnalizeChartData {
  name: string;
  Норма: number;
  Патологія: number;
}

export interface IImage {
  date: string;
  imageId: number;
  link: string;
  name: string;
  patientId: number;
  type: string;
}

export interface IPatient {
  age: number;
  diagnosis: string;
  doctorId: string;
  fathersName: string;
  firstName: string;
  height: number;
  lastName: string;
  patientId: string;
  weight: number;
}

export type IAppProps = {
  children?: React.ReactNode;
};
