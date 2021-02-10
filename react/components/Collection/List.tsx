import React, { useState } from 'react';
import useSWR from 'swr';
import SubcategoryList from '@/components/Category/SubcategoryList';

interface Category {
  id: string;
  name: string;
  is_active: boolean;
}

const CategoryList: React.FC = () => {
  const { data: collections } = useSWR(
    `${process.env.GO_MONGO_ENDPOINT_API}/collection-level-1`
  );
  //   const [showSubCategory, setShowSubCategory] = useState<boolean>(false);

  //   const toggleSubCategory = () => {
  //     setShowSubCategory(true);
  //   };

  return (
    <React.Fragment>
      {collections?.data.map((collection) => (
        <div key={collection._id}>
          <h3>{collection.name}</h3>
          {collection.sub_collection.map((sub) => (
            <div key={sub._id}>
              <h4>- {sub.name}</h4>
              <ul>
                {Object.keys(sub.collection_level_3).map((index) => (
                  <li key={index}>{sub.collection_level_3[index].name}</li>
                ))}
              </ul>
            </div>
          ))}
        </div>
      ))}
    </React.Fragment>
  );
};

export default CategoryList;
