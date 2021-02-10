import React from 'react';
import Link from 'next/link';

const NavGlobal: React.FC = () => {
    return (
        <React.Fragment>
            <span><Link href="/product"><a>Product</a></Link></span>
            <span><Link href="/collection
            "><a>Collection</a></Link></span>
        </React.Fragment>
    )
}

export default NavGlobal