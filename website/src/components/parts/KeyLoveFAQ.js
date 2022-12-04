import React, { useEffect, useRef } from "react";
import { Icon } from "@iconify/react";

const KeyFeaturesData = [
  {
    icon: "share-2",
    content: "Comprehensive Workflows",
  },
  {
    icon: "zap",
    content: "High Throughput",
  },
  {
    icon: "activity",
    content: "Fault Tolerance",
  },
  {
    icon: "lock",
    content: "End-to-End Encryption",
  },
  {
    icon: "copy",
    content: "Deduplication",
  },
  {
    icon: "minimize-2",
    content: "Compression",
  },
  {
    icon: "code",
    content: "Open Source",
  },
  {
    icon: "upload-cloud",
    content: "Choice of Object Storage",
  },
  {
    icon: "check-circle",
    content: "Retention Policies",
  },
];

const AccordionItemsData = [
  {
    id: "One",
    title: "What platforms does Corso run on?",
    description:
      "Corso has both native binaries and container images for Windows, Linux, and macOS.",
  },
  {
    id: "Two",
    title: "What Microsoft 365 services can I backup using Corso?",
    description:
      "Corso currently supports OneDrive and Exchange. Support for Teams and SharePoint is in active development and is therefore not recommended for production use.",
  },
  {
    id: "Three",
    title: "What object storage does Corso support?",
    description:
      "Corso supports any S3-compliant object storage system including AWS S3 (including Glacier Instant Access), Google Cloud Storage, and Backblaze. Azure Blob support is coming soon.",
  },
  {
    id: "Four",
    title: "How can I get help for Corso?",
    comp: (
      <>
        {" "}
        If you are unable to find an answer in our documentation, please file{" "}
        <a
          href="https://github.com/alcionai/corso/issues"
          className="text-indigo-600"
          target="_blank"
        >
          GitHub issues
        </a>{" "}
        for bugs or join the{" "}
        <a
          href="https://discord.gg/63DTTSnuhT"
          className="text-indigo-600"
          target="_blank"
        >
          Discord community
        </a>
        .
      </>
    ),
  },
  {
    id: "Five",
    title: "What is Corso's open-source license?",
    description:
      "Corso's source code is licensed under the OSI-approved Apache v2 open-source license.",
  },
  {
    id: "Six",
    title: "How do I request a new feature?",
    comp: (
      <>
        {" "}
        You can request new features by creating a{" "}
        <a
          href="https://github.com/alcionai/corso/issues"
          className="text-indigo-600"
          target="_blank"
        >
          new GitHub issue
        </a>{" "}
        and labeling it as an enhancement.
      </>
    ),
  },
];

export default function KeyLoveFAQ() {
  const jarallaxRef = useRef(null);
  useEffect(() => {
    if (typeof window !== "undefined") {
      const WOW = require("wowjs");
      const father = require("feather-icons");
      const jarallax = require("jarallax");
      require("tw-elements");

      new WOW.WOW({
        live: false,
      }).init();
      father.replace();
      jarallax.jarallax(jarallaxRef.current, {
        speed: 0.2,
      });
    }
  }, []);

  return (
    <section className="relative third-section---custom md:py-24 !tracking-wide py-16 overflow-hidden">
      <div className="container">
        <div
          className="grid grid-cols-1 pb-8 text-center wow animate__animated animate__fadeInUp"
          data-wow-delay=".1s"
        >
          <h3
            className={`mb-6 h3-1---custom mt-8 md:text-4xl text-3xl md:leading-normal leading-normal font-bold`}
          >
            Key Features
          </h3>

          <p className={`p-1---custom max-w-xl mx-auto`}>
            See why Corso is a perfect fit for your Microsoft 365 backup and
            recovery needs.
          </p>
        </div>

        <div className="grid lg:grid-cols-3 md:grid-cols-2 grid-cols-1 grid-flow-row-dense gap-[30px] mt-8">
          {KeyFeaturesData.map((item, index) => (
            <KeyFeaturesCard key={index} {...item} />
          ))}
        </div>
      </div>

      <div className="container md:mt-24 mt-16">
        <div className="container lg mx-auto">
          <div className="grid grid-cols-1 pb-2 text-center wow animate__animated animate__fadeInUp">
            <h3 className="mb-6 mt-8 md:text-4xl text-3xl md:leading-normal leading-normal font-bold">
              Why Everyone{" "}
              <span className="after:absolute after:right-0 after:left-0 after:bottom-1 after:lg:h-3 after:h-2 after:w-auto after:rounded-md after:bg-indigo-600/30 relative text-indigo-600">
                Loves
                <div className="absolute right-0 left-0 bottom-1 lg:h-3 h-2 w-auto rounded-md bg-indigo-600/30"></div>
              </span>{" "}
              Corso
            </h3>
          </div>
        </div>

        <div className="grid md:grid-cols-2 grid-cols-1 items-center gap-[30px]">
          <div
            className="relative wow animate__animated animate__fadeInLeft"
            data-wow-delay=".3s"
          >
            <img
              src="assets/images/why/chat.svg"
              className="rounded-lg"
              alt="Group discussion"
            />
            <div className="overflow-hidden absolute lg:h-[400px] h-[320px] lg:w-[400px] w-[320px] bg-indigo-600/5 bottom-0 left-0 rotate-45 -z-1 rounded-3xl"></div>
          </div>

          <div
            className="lg:ml-8 wow animate__animated animate__fadeInRight"
            data-wow-delay=".3s"
          >
            <h3 className="mb-4 text-3xl leading-normal font-medium">
              Community
            </h3>
            <p className="text-slate-400">
              The Corso community provides a venue for M365 admins to share and
              learn about the importance of data protection as well as best
              practices around M365 secure configuration and compliance
              management.
            </p>
            <ul className="list-none text-slate-400 mt-4">
              <li className="mb-1 flex">
                <Icon
                  className="text-indigo-600 text-xl mr-2"
                  icon="material-symbols:check-circle-outline"
                />{" "}
                Community-led blogs, forums, and discussions
              </li>
              <li className="mb-1 flex">
                <Icon
                  className="text-indigo-600 text-xl mr-2"
                  icon="material-symbols:check-circle-outline"
                />{" "}
                Public and feedback-driven development roadmap{" "}
              </li>
              <li className="mb-1 flex">
                <Icon
                  className="text-indigo-600 text-xl mr-2"
                  icon="material-symbols:check-circle-outline"
                />{" "}
                All community contributions welcome
              </li>
            </ul>

            <div className="mt-4">
              <a
                href="https://discord.gg/63DTTSnuhT"
                target="_blank"
                className="btn btn-link !no-underline link-underline link-underline-black text-indigo-600 hover:text-indigo-600 after:bg-indigo-600 duration-500 ease-in-out"
              >
                Join Us On Discord{" "}
                <Icon icon="uim:angle-right-b" className="align-middle" />
              </a>
            </div>
          </div>
        </div>
      </div>

      <div className="container md:mt-24 mt-16">
        <div className="grid md:grid-cols-2 grid-cols-1 items-center gap-[30px]">
          <div
            className="relative order-1 md:order-2 wow animate__animated animate__fadeInRight"
            data-wow-delay=".5s"
          >
            <img
              src="assets/images/why/security.svg"
              className="rounded-lg"
              alt="Approval of fingerprint security"
            />
            <div className="overflow-hidden absolute lg:h-[400px] h-[320px] lg:w-[400px] w-[320px] bg-indigo-600/5 bottom-0 right-0 rotate-45 -z-1 rounded-3xl"></div>
          </div>

          <div
            className="lg:mr-8 order-2 md:order-1 wow animate__animated animate__fadeInLeft"
            data-wow-delay=".5s"
          >
            <h3 className="mb-4 text-3xl leading-normal font-medium">
              Data Security
            </h3>
            <p className="text-slate-400">
              Corso provides secure data backup that protects customers against
              accidental data loss, service provider downtime, and malicious
              threats including ransomware attacks.
            </p>
            <ul className="list-none text-slate-400 mt-4">
              <li className="mb-1 flex">
                <Icon
                  className="text-indigo-600 text-xl mr-2"
                  icon="material-symbols:check-circle-outline"
                />{" "}
                End-to-end zero-trust AES-256 and TLS encryption
              </li>
              <li className="mb-1 flex">
                <Icon
                  className="text-indigo-600 text-xl mr-2"
                  icon="material-symbols:check-circle-outline"
                />{" "}
                Support for air-gapped backup storage
              </li>
              <li className="mb-1 flex">
                <Icon
                  className="text-indigo-600 text-xl mr-2"
                  icon="material-symbols:check-circle-outline"
                />{" "}
                Choice of backup storage provider and geo location
              </li>
            </ul>
          </div>
        </div>
      </div>

      <div className="container md:mt-24 mt-16">
        <div className="grid md:grid-cols-2 grid-cols-1 items-center mt-8 gap-[30px]">
          <div
            className="relative wow animate__animated animate__fadeInLeft"
            data-wow-delay=".5s"
          >
            <img
              src="assets/images/why/data.svg"
              className="rounded-lg"
              alt="Data extraction dashboard"
            />
            <div className="overflow-hidden absolute lg:h-[400px] h-[320px] lg:w-[400px] w-[320px] bg-indigo-600/5 bottom-0 left-0 rotate-45 -z-1 rounded-3xl"></div>
          </div>

          <div
            className="lg:ml-8 wow animate__animated animate__fadeInRight"
            data-wow-delay=".5s"
          >
            <h3 className="mb-4 text-3xl leading-normal font-medium">
              Robust Backups
            </h3>
            <p className="text-slate-400">
              Corso, purpose-built for M365 protection, provides easy-to-use
              comprehensive backup and restore workflows that reduces backup
              time, improve time-to-recovery, reduce admin overhead, and replace
              unreliable scripts or workarounds.
            </p>
            <ul className="list-none text-slate-400 mt-4">
              <li className="mb-1 flex">
                <Icon
                  className="text-indigo-600 text-xl mr-2"
                  icon="material-symbols:check-circle-outline"
                />{" "}
                Constantly updated M365 Graph Data engine
              </li>
              <li className="mb-1 flex">
                <Icon
                  className="text-indigo-600 text-xl mr-2"
                  icon="material-symbols:check-circle-outline"
                />{" "}
                Purpose-built, flexible, fine-grained data protection workflows
              </li>
              <li className="mb-1 flex">
                <Icon
                  className="text-indigo-600 text-xl mr-2"
                  icon="material-symbols:check-circle-outline"
                />{" "}
                High-performance backup and recovery data movers
              </li>
            </ul>

            <div className="mt-4">
              <a
                href="https://docs.corsobackup.io/quickstart"
                target="_blank"
                className="btn btn-link !no-underline link-underline link-underline-black text-indigo-600 hover:text-indigo-600 after:bg-indigo-600 duration-500 ease-in-out"
              >
                Use The Quick Start For Your First Backup{" "}
                <Icon icon="uim:angle-right-b" className="align-middle" />
              </a>
            </div>
          </div>
        </div>
      </div>

      <div className="container md:mt-24 mt-16">
        <div className="grid md:grid-cols-2 grid-cols-1 items-center gap-[30px]">
          <div
            className="relative order-1 md:order-2 wow animate__animated animate__fadeInRight"
            data-wow-delay=".5s"
          >
            <img
              src="assets/images/why/savings.svg"
              className="rounded-lg"
              alt="Adding money to a savings jar"
            />
            <div className="overflow-hidden absolute lg:h-[400px] h-[320px] lg:w-[400px] w-[320px] bg-indigo-600/5 bottom-0 right-0 rotate-45 -z-1 rounded-3xl"></div>
          </div>

          <div
            className="lg:mr-8 order-2 md:order-1 wow animate__animated animate__fadeInLeft"
            data-wow-delay=".5s"
          >
            <h3 className="mb-4 text-3xl leading-normal font-medium">
              Cost Savings
            </h3>
            <p className="text-slate-400">
              Corso, a 100% open-source tool, provides a free alternative for
              cost-conscious teams. It further reduces storage costs by
              supporting flexible retention policies and efficiently compressing
              and deduplicating data before storing it in low-cost cloud object
              storage.
            </p>
            <ul className="list-none text-slate-400 mt-4">
              <li className="mb-1 flex">
                <Icon
                  className="text-indigo-600 text-xl mr-2"
                  icon="material-symbols:check-circle-outline"
                />{" "}
                Free forever OSS with no licensing costs
              </li>
              <li className="mb-1 flex">
                <Icon
                  className="text-indigo-600 text-xl mr-2"
                  icon="material-symbols:check-circle-outline"
                />{" "}
                Client-side compression and deduplication
              </li>
              <li className="mb-1 flex">
                <Icon
                  className="text-indigo-600 text-xl mr-2"
                  icon="material-symbols:check-circle-outline"
                />{" "}
                Support for S3-compliant storage including AWS Glacier IA
              </li>
            </ul>

            <div className="mt-4">
              <a
                href="https://docs.corsobackup.io/setup/repos"
                target="_blank"
                className="btn btn-link !no-underline link-underline link-underline-black text-indigo-600 hover:text-indigo-600 after:bg-indigo-600 duration-500 ease-in-out"
              >
                Read about our Object Storage support{" "}
                <Icon icon="uim:angle-right-b" className="align-middle" />
              </a>
            </div>
          </div>
        </div>
      </div>

      {/* Accordions */}

      <div className="container md:mb-8 mb-4 md:mt-24 mt-16 wow animate__animated animate__fadeInUp">
        <div className="grid grid-cols-1 pb-8 text-center">
          <h3
            className={`mb-6 mt-8 h3-1---custom md:text-4xl text-3xl md:leading-normal leading-normal font-bold`}
          >
            Frequently Asked Questions
          </h3>
        </div>

        <div className="relative grid md:grid-cols-12 grid-cols-1 items-center gap-[30px]">
          <div className="md:col-span-6">
            <div className="relative">
              <div className="relative rounded-xl overflow-hidden shadow-md dark:shadow-gray-800">
                <div
                  ref={jarallaxRef}
                  className="w-full jarallax py-72 bg-slate-400 custom-bg_ bg-no-repeat bg-top"
                  data-jarallax='{"speed": 0.1}'
                ></div>
              </div>
            </div>
          </div>

          <div className="md:col-span-6">
            <div className="accordion space-y-3" id="accordionExample">
              {AccordionItemsData.map((item, i) => (
                <AccordionItemCard {...item} key={i} />
              ))}
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}

const AccordionItemCard = ({ title, description, id, comp }) => {
  return (
    <div className="accordion-item text-white relative  shadow shadow-gray-800 rounded-md overflow-hidden">
      <h2
        className="accordion-header mb-0 !cursor-pointer font-semibold"
        id={`heading${id}`}
      >
        <button
          className={`transition accordion-button-custom h3-1---custom !text-base !cursor-pointer border-none outline-none collapsed focus:outline-none !bg-transparent flex justify-between items-center p-5 w-full font-medium text-left`}
          type="button"
          data-bs-toggle="collapse"
          data-bs-target={`#collapse${id}`}
          aria-expanded="false"
          aria-controls={`collapse${id}`}
        >
          <span>{title}</span>
        </button>
      </h2>
      <div
        id={`collapse${id}`}
        className="accordion-collapse collapse"
        aria-labelledby={`heading${id}`}
        data-bs-parent="#accordionExample"
      >
        <div className="accordion-body p-5">
          <p className="visible text-gray-400">{comp ? comp : description}</p>
        </div>
      </div>
    </div>
  );
};

const KeyFeaturesCard = ({ content, icon }) => {
  return (
    <div
      className="wow animate__animated animate__fadeInUp"
      data-wow-delay=".1s"
    >
      <div
        className={`flex transition-all duration-500 scale-hover shadow shadow-gray-800 hover:shadow-md hover:shadow-gray-700 ease-in-out items-center p-3 rounded-md 
        div-1---custom
        `}
      >
        <div className="flex items-center justify-center h-[45px] min-w-[45px] -rotate-45 bg-gradient-to-r from-transparent to-indigo-600/10 text-indigo-600 text-center rounded-full mr-3">
          <i data-feather={icon} className="h-5 w-5 rotate-45"></i>
        </div>
        <div className="flex-1">
          <h4 className="mb-0 text-lg font-bold">{content}</h4>
        </div>
      </div>
    </div>
  );
};
