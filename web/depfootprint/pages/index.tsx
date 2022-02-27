import { useRef } from 'react';
import type { NextPage } from 'next'
import { useRouter } from 'next/router';

import BaseLayout from '../components/BaseLayout';

const Home: NextPage = () => {
  const inputEl = useRef<HTMLInputElement>(null);
  const router = useRouter();

  function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const value = inputEl?.current?.value;
    if (!value) return;
    router.push(`/packages/${value}`);
  }

  return (
    <BaseLayout>
      <div className="container mx-auto my-5">
        <form onSubmit={handleSubmit}>
          <label>
            NPM package:
            <input placeholder="express" ref={inputEl} className="bg-gray-200 appearance-none border-2 border-gray-200 rounded py-2 px-1 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-blue-500 mx-2" type="text" />
          </label>
          <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">Go</button>
        </form>
      </div>
    </BaseLayout>
    )
}

export default Home
