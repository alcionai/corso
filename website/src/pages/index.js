import React, { useEffect } from "react";
import Layout from "@theme/Layout";
import Head from "@docusaurus/Head"
import { MainComp } from "@site/src/components/parts/MainComp";
import { useColorMode } from "@docusaurus/theme-common";

const ThemeColor = () => {
  const { colorMode, setColorMode } = useColorMode();

  useEffect(function () {
    if (window.location.pathname === "/") {
      if (colorMode !== "dark") {
        //force dark theme to home page without overriding user settings
        setColorMode("dark", { persist: false });
      }
    } else {
      setColorMode(localStorage.getItem("theme"));
    }
  });

  return null;
};

export default function Home() {
  return (
    <Layout
      title="Free, Secure, and Open-Source Backup for Microsoft 365"
      description="Intro, docs, and blog for Corso, an open-source tool, that protects Microsoft 365 data by securely and efficiently backing up all business-critical data to object storage.">
       <Head>
        <script type="application/ld+json">{`
        {
          "@context" : "https://schema.org",
          "@type" : "WebSite",
          "name" : "Corso",
          "url" : "https://corsobackup.io/"
        }
        `}</script>
      </Head>
      <ThemeColor />
      <MainComp />
    </Layout>
  );
}
