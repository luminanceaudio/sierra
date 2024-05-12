import { useState } from 'react';
import Layout from '../../components/Layout/Layout';
import Widget from '../../components/Widget/Widget';
import mountainSignImg from '../../../../assets/mountain_sign.svg';
import Title from '../../components/Title/Title';
import Button from '../../components/Button/Button';
import { useSources } from '../../../api/sources';
import { Source, SourceNew } from './Source';

function Sources() {
  const { data: sources } = useSources();

  return (
    <Layout title="Sources">
      <Widget
        style={{
          display: 'flex',
          flex: 1,
          alignItems: 'center',
          justifyContent: 'center',
          paddingBottom: 100,
        }}
      >
        <div
          style={{
            display: 'flex',
            flex: 1,
            flexDirection: 'column',
            gap: 15,
          }}
        >
          <SourceNew />
          {sources?.data?.sources.map((source) => (
            <Source key={source.uri} uri={source.uri} />
          ))}
        </div>
      </Widget>
    </Layout>
  );
}

export default Sources;
