import React from "react";
import Layout from "@theme/Layout";
import { MainComp } from "../components/parts/MainComp";

export default function Home() {
  return (
    <Layout
      title="Home Page"
      description="Documentation for Corso, a free, secure, and open-source backup tool for Microsoft 365"
    >
      <MainComp />
    </Layout>
  );
}
