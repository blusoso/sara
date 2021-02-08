import React from 'react';
import Link from 'next/link';

const NavGlobal: React.FC = () => {
    return (
        <React.Fragment>
            <span><Link href="/product"><a>Product</a></Link></span>
            <span><Link href="/category"><a>Category</a></Link></span>
        </React.Fragment>
    )
}

export default NavGlobal