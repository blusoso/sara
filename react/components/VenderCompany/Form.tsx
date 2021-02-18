import React from 'react';
import { useForm } from 'react-hook-form';

interface VenderCompany {
  id: string;
  name: string;
  address: string;
  phone_number: string;
  created_at: string;
  updated_at: string;
  deleted_at?: string;
}

interface Props {
  submitVenderCompany: (venderCompany: VenderCompany) => void;
  defaultVenderCompany?: VenderCompany;
}

const VenderCompanyForm: React.FC<Props> = ({
  submitVenderCompany,
  defaultVenderCompany,
}) => {
  const { register, handleSubmit, errors } = useForm({
    defaultValues: {
      name: defaultVenderCompany?.name,
      address: defaultVenderCompany?.address,
      phone_number: defaultVenderCompany?.phone_number,
    },
  });

  const onSubmit = (venderCompany: VenderCompany) => {
    submitVenderCompany(venderCompany);
  };

  return (
    <React.Fragment>
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className="form-group">
          <label>Name</label>
          <input
            type="text"
            name="name"
            ref={register({ required: 'Please enter a name' })}
          />
          {errors.name?.message}
        </div>
        <div className="form-group">
          <label>Address</label>
          <input
            type="text"
            name="address"
            ref={register({ required: 'Please enter a address' })}
          />
          {errors.address?.message}
        </div>
        <div className="form-group">
          <label>Phone number</label>
          <input
            type="text"
            name="phone_number"
            ref={register({ required: 'Please enter a Phone number' })}
          />
          {errors.phone_number?.message}
        </div>
        <div>
          <button type="submit">Submit</button>
        </div>
      </form>
    </React.Fragment>
  );
};

export default VenderCompanyForm;
