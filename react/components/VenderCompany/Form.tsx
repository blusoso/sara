import React from 'react';
import { useForm } from 'react-hook-form';

const VenderCompanyForm: React.FC = ({
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

  const onSubmit = (venderCompany) => submitVenderCompany(venderCompany);

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
