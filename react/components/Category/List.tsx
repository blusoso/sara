import React, { useState } from 'react';
import useSWR from 'swr';
import SubcategoryList from '@/components/Category/SubcategoryList';

interface Category {
  id: string;
  name: string;
  is_active: boolean;
}

const CategoryList: React.FC = () => {
  const { data: categories } = useSWR(
    `${process.env.GO_MONGO_ENDPOINT_API}/category`
  );
  const [showSubCategory, setShowSubCategory] = useState<boolean>(false);

  const toggleSubCategory = () => {
    setShowSubCategory(true);
  };

  return (
    <React.Fragment>
      {categories.map((category) => (
        <div key={category.id}>
          <div onClick={toggleSubCategory}>{category.name}</div>
          {showSubCategory && <div>Test</div>}
        </div>
      ))}
    </React.Fragment>
  );
};

export default CategoryList;
